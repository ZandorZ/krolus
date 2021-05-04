package feed

import (
	"krolus/feed/providers"
	"krolus/models"
	"sync"

	"github.com/mmcdole/gofeed"
)

// FacadeChecker A checker that uses gofeed.Parser and a Proxy of Providers
type FacadeChecker struct {
	Requester
	*gofeed.Parser
	proxy          *providers.Proxy
	feedParserPool sync.Pool
}

// request and parse
func (c *FacadeChecker) request(url string) (*gofeed.Feed, error) {
	response, err := c.Request(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	fp := c.feedParserPool.Get().(*gofeed.Parser)
	f, err := fp.Parse(response.Body)
	if err != nil {
		return nil, err
	}
	c.feedParserPool.Put(fp)

	return f, nil
}

// Check Checks a sub and return collection of new items
func (c *FacadeChecker) Check(sub *models.SubscriptionModel) (models.ItemCollection, error) {

	f, err := c.request(sub.XURL)
	if err != nil {
		return nil, err
	}

	return c.proxy.GetNewItems(sub, f), nil
}

// NewChecker ...
func NewChecker(req Requester) Checker {
	return &FacadeChecker{
		Requester: req,
		Parser:    gofeed.NewParser(),
		proxy:     providers.NewProxy(),
		feedParserPool: sync.Pool{
			New: func() interface{} {
				return gofeed.NewParser()
			},
		},
	}
}
