package providers

import (
	"krolus/models"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
)

var redditPatcher PatchFunc = func(item *models.ItemModel) {

}

var redditConverter ConvertFunc = func(sub *models.SubscriptionModel, feed *gofeed.Feed) models.ItemCollection {

	items := make(models.ItemCollection, len(feed.Items))

	sub.Provider = "reddit"
	for i, item := range feed.Items {
		items[i] = models.ItemModel{
			ID:           uuid.New().String(),
			Title:        item.Title,
			Link:         item.Link,
			Description:  item.Description,
			Content:      item.Content,
			New:          true,
			Thumbnail:    getThumbReddit(item),
			Published:    item.PublishedParsed.Local(),
			Provider:     "reddit",
			Type:         "unknown",
			Subscription: sub.ID,
		}
	}
	return items
}

func getThumbReddit(item *gofeed.Item) string {
	if item.Extensions != nil {
		return item.Extensions["media"]["thumbnail"][0].Attrs["url"]
	}
	return ""
}
