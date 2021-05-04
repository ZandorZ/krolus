package providers

import (
	"fmt"
	"krolus/models"
	"regexp"
	"strings"

	"github.com/mmcdole/gofeed"
)

const YOUTUBE = "youtube"

type YoutubeProvider struct {
	*Proxy
}

func NewYoutubeProvider(p *Proxy) Provider {
	return &YoutubeProvider{
		Proxy: p,
	}
}

func (p *YoutubeProvider) Convert(item *gofeed.Item, model *models.ItemModel) {
	model.Description = p.getDescription(item)
	model.Content = ""
	model.Thumbnail = p.getThumbnail(item)
	model.Provider = YOUTUBE
	model.Type = models.TypeVideo
}

func (p *YoutubeProvider) Fetch(item *models.ItemModel) {
	id, _ := p.extractID(item.Link)
	item.Type = models.TypeVideo
	item.Content = fmt.Sprintf("https://www.youtube.com/embed/%s?ecver=1&amp;autoplay=1&amp;iv_load_policy=3&amp;rel=0&amp;yt:stretch=1:1&amp;autohide=1&amp;color=red", id)
}

func (p *YoutubeProvider) Download(item *models.ItemModel) {

}

func (p *YoutubeProvider) extractID(content string) (string, bool) {
	re, err := regexp.Compile(`^.*((youtu.be\/)|(v\/)|(\/u\/\w\/)|(embed\/)|(watch\?))\??v?=?([^#\&\?]*).*`)

	if err != nil {
		return "", false
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
