package providers

import (
	"fmt"
	"krolus/models"
	"regexp"

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
			Description:  item.Extensions["media"]["group"][0].Children["description"][0].Value,
			Content:      "",
			New:          true,
			Thumbnail:    item.Extensions["media"]["group"][0].Children["thumbnail"][0].Attrs["url"],
			Published:    item.PublishedParsed.Local(),
			Provider:     "youtube",
			Type:         "video",
			Subscription: sub.ID,
		}
	}
	return items
}

var youtubeFetcher FetcherFunc = func(item *models.ItemModel) {
	id, _ := extractYoutube(item.Link)
	item.Content = fmt.Sprintf("https://www.youtube.com/embed/%s?ecver=1&amp;iv_load_policy=3&amp;rel=0&amp;showinfo=0&amp;yt:stretch=16:9&amp;autohide=1&amp;color=red&amp;width=560&amp;width=560&amp;version=3&amp;vq=hd720", id)
}

func extractYoutube(content string) (string, bool) {
	re, err := regexp.Compile(`^.*((youtu.be\/)|(v\/)|(\/u\/\w\/)|(embed\/)|(watch\?))\??v?=?([^#\&\?]*).*`)

	if err != nil {
		panic(err)
	}
	found := re.FindAllStringSubmatch(content, 2)

	if len(found) > 0 {
		return found[0][7], true
	}
	return "", false
}
