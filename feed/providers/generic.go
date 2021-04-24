package providers

import (
	"krolus/models"
	"time"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
)

var genericConverter ConvertFunc = func(sub *models.SubscriptionModel, feed *gofeed.Feed) models.ItemCollection {

	items := make(models.ItemCollection, len(feed.Items))

	sub.Provider = "generic"
	for i, item := range feed.Items {
		items[i] = models.ItemModel{
			ID:           uuid.New().String(),
			Title:        item.Title,
			Link:         item.Link,
			Description:  item.Description,
			Content:      item.Content,
			New:          true,
			Thumbnail:    getImage(item),
			Published:    getDate(item),
			Provider:     "generic",
			Type:         "unknown",
			Subscription: sub.ID,
		}
	}
	return items
}

var genericPatcher PatchFunc = func(item *models.ItemModel) {

}

func getImage(item *gofeed.Item) string {
	img := ""
	if item.Image != nil {
		img = item.Image.URL
	} else if len(item.Enclosures) > 0 && item.Enclosures[0].Type == "image/jpeg" {
		img = item.Enclosures[0].URL
	}
	return img
}

func getDate(item *gofeed.Item) time.Time {
	d := item.PublishedParsed
	if d == nil {
		d = item.UpdatedParsed
	}
	if d == nil {
		d = &time.Time{}
	}
	return d.Local()
}
