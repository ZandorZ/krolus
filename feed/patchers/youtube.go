package patchers

import (
	"fmt"
	"krolus/models"
	"regexp"

	"github.com/mmcdole/gofeed"
)

// Youtube ...
type Youtube struct {
}

// NewYoutubePatcher ....
func NewYoutubePatcher() Patcher {
	return &Youtube{}
}

// Patch ...
func (yt *Youtube) Patch(item *gofeed.Item) *models.ItemModel {

	id := item.Extensions["yt"]["videoId"][0].Value

	return &models.ItemModel{
		ID:          id,
		Title:       item.Title,
		Link:        item.Link,
		Description: item.Extensions["media"]["group"][0].Children["description"][0].Value,
		Thumbnail:   item.Extensions["media"]["group"][0].Children["thumbnail"][0].Attrs["url"],
		Published:   item.PublishedParsed.Local(),
		New:         true,
		Provider:    "youtube",
		Type:        "video",
		Embed:       embedYoutube(id),
	}
}

func extractYoutube(content string) (string, bool) {
	re, err := regexp.Compile(`(?:youtube(?:-nocookie)?\.com\/(?:[^\/\n\s]+\/\S+\/|(?:v|e(?:mbed)?)\/|\S*?[?&]v=)|youtu\.be\/)([a-zA-Z0-9_-]{11})\W`)

	if err != nil {
		panic(err)
	}

	found := re.FindAllStringSubmatch(content, 2)

	if len(found) > 0 {
		return found[0][1], true
	}

	return "", false
}

func embedYoutube(id string) string {
	return fmt.Sprintf("https://www.youtube.com/embed/%s?ecver=1&amp;iv_load_policy=3&amp;rel=0&amp;showinfo=0&amp;yt:stretch=16:9&amp;autohide=1&amp;color=red&amp;width=560&amp;width=560&amp;version=3&amp;vq=hd720", id)
}
