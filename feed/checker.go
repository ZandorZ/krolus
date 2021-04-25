package feed

import (
	"krolus/feed/providers"
	"krolus/models"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/thoas/go-funk"
)

// FacadeChecker A checker that uses gofeed.Parser and a Proxy of Providers
type FacadeChecker struct {
	Requester
	*gofeed.Parser
}

func (c *FacadeChecker) request(url string) (*gofeed.Feed, error) {
	response, err := c.Request(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	f, err := c.Parse(response.Body)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// newItems only new items
func (c *FacadeChecker) newItems(sub *models.SubscriptionModel, f *gofeed.Feed) {

	var lastUpdate time.Time
	var lastItemLink string

	items := funk.Filter(f.Items, func(i *gofeed.Item) bool {

		if i.Link != sub.LastItemLink && i.PublishedParsed.After(sub.LastUpdate) {
			if i.PublishedParsed.After(lastUpdate) {
				lastUpdate = *i.PublishedParsed
				lastItemLink = i.Link
			}
			return true
		}
		return false
	}).([]*gofeed.Item)

	f.Items = items
	if lastItemLink != "" {
		sub.LastUpdate = lastUpdate
		sub.LastItemLink = lastItemLink
	}
}

// Check Checks a sub and return collection of new items
func (c *FacadeChecker) Check(sub *models.SubscriptionModel) (models.ItemCollection, error) {

	f, err := c.request(sub.XURL)
	if err != nil {
		return nil, err
	}
	c.newItems(sub, f)

	return providers.NewProxy(sub.XURL).Convert(sub, f), nil
}

func NewChecker(req Requester) Checker {
	return &FacadeChecker{
		Requester: req,
		Parser:    gofeed.NewParser(),
	}
}
