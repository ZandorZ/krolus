package providers

import (
	"html"
	"krolus/models"
	"net/url"
	"regexp"

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

var redditFetcher FetcherFunc = func(proxy *Proxy, item *models.ItemModel) {
	item.Description = item.Content
	item.Link = extractLinkReddit(item.Content)
	u, err := url.Parse(item.Link)
	if err != nil || u.Hostname() == "reddit.com" || u.Hostname() == "www.reddit.com" {
		return
	}
	proxy.Fetch(item)
}

func getThumbReddit(item *gofeed.Item) string {
	if item.Extensions != nil {
		return item.Extensions["media"]["thumbnail"][0].Attrs["url"]
	}
	return ""
}

func extractLinkReddit(content string) string {

	content = html.UnescapeString(content)
	// pattern := `<a href="(((https?\:\/\/)|(www\.))(\S+))">\[link\]</a>`
	pattern := `<a href\s*=\s*["\']?([^"\'\s>]+)["\']?>\[link\]`
	re, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}

	found := re.FindAllStringSubmatch(content, 2)
	if len(found) > 0 {
		return found[0][1]
	}

	return ""
}
