package providers

import (
	"krolus/models"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
)

var youtubePatcher PatchFunc = func(item *models.ItemModel) {
	item.Type = models.TypeVideo
}

var yotubeConverter ConvertFunc = func(sub *models.SubscriptionModel, feed *gofeed.Feed) models.ItemCollection {

	items := make(models.ItemCollection, len(feed.Items))

	sub.Provider = "youtube"
	for i, item := range feed.Items {
		items[i] = models.ItemModel{
			ID:           uuid.New().String(),
			Title:        item.Title,
			Link:         item.Link,
			Description:  item.Description,
			Content:      item.Content,
			New:          true,
			Thumbnail:    getThumbYoutube(item),
			Published:    item.PublishedParsed.Local(),
			Provider:     "youtube",
			Type:         "unknown",
			Subscription: sub.ID,
		}
	}
	return items
}

func getThumbYoutube(item *gofeed.Item) string {
	if item.Extensions != nil {
		return item.Extensions["media"]["group"][0].Children["thumbnail"][0].Attrs["url"]
	}
	return ""
}
