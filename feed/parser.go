package feed

import (
	"krolus/feed/patchers"
	"krolus/models"
	"time"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
)

// GenericParser ...
type GenericParser struct {
	requester  Requester
	feedParser *gofeed.Parser
}

// NewParser ...
func NewParser(requester Requester) Parser {

	return &GenericParser{
		requester:  requester,
		feedParser: gofeed.NewParser(),
	}
}

// Parse ...
func (p *GenericParser) Parse(sub *models.SubscriptionModel) (models.ItemCollection, error) {

	var items models.ItemCollection

	// TODO: encapsulate request ///////////////////
	// Make request
	response, err := p.requester.Request(sub.XURL)
	if err != nil {
		return items, err
	}
	defer response.Body.Close()
	///////////////////////////////////////////////

	feed, err := p.feedParser.Parse(response.Body)
	if err != nil {
		return nil, err
	}

	var lastUpdate time.Time
	var lastItem string
	for _, item := range feed.Items {

		newItem := patchers.GetPatcher(item.Link)(item)
		newItem.Subscription = sub.ID
		newItem.ID = uuid.New().String()

		if item.Link != sub.LastItem && newItem.Published.After(sub.LastUpdate) {
			if newItem.Published.After(lastUpdate) {
				lastUpdate = newItem.Published
			}
			items = append(items, *newItem)
			lastItem = item.Link
		}
	}

	if lastItem != "" {
		sub.LastUpdate = lastUpdate
		sub.LastItem = lastItem
	}

	return items, nil
}
