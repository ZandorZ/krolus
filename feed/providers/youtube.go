package providers

import (
	"fmt"
	"krolus/models"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
)

type YoutubeProvider struct {
	*Proxy
}

func NewYoutubeProvider(p *Proxy) Provider {
	return &YoutubeProvider{
		Proxy: p,
	}
}

func (p *YoutubeProvider) Convert(item *gofeed.Item) *models.ItemModel {
	return &models.ItemModel{
		ID:          uuid.New().String(),
		Title:       item.Title,
		Link:        item.Link,
		Description: p.getDescription(item),
		Content:     "",
		New:         true,
		Thumbnail:   p.getThumbnail(item),
		Published:   getDate(item),
		Provider:    "youtube",
		Type:        models.TypeVideo,
	}
}

func (p *YoutubeProvider) Fetch(item *models.ItemModel) {
	id, _ := p.extractID(item.Link)
	item.Type = models.TypeVideo
	item.Content = fmt.Sprintf("https://www.youtube.com/embed/%s?ecver=1&amp;iv_load_policy=3&amp;rel=0&amp;showinfo=0&amp;yt:stretch=16:9&amp;autohide=1&amp;color=red&amp;width=560&amp;width=560&amp;version=3&amp;vq=hd720", id)
}

func (p *YoutubeProvider) Download(item *models.ItemModel) error {
	return nil
}

func (p *YoutubeProvider) extractID(content string) (string, bool) {
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

func (p *YoutubeProvider) getDescription(item *gofeed.Item) string {
	//sub not youtube
	if !strings.HasPrefix(item.GUID, "yt:video:") {
		return item.Description
	}
	return item.Extensions["media"]["group"][0].Children["description"][0].Value
}

func (p *YoutubeProvider) getThumbnail(item *gofeed.Item) string {
	//sub not youtube
	if !strings.HasPrefix(item.GUID, "yt:video:") {
		id, _ := p.extractID(item.Link)
		return fmt.Sprintf("https://i.ytimg.com/vi/%s/mqdefault.jpg", id)
	}
	return item.Extensions["media"]["group"][0].Children["thumbnail"][0].Attrs["url"]
}
