package providers

import (
	"html"
	"krolus/models"
	"regexp"

	"github.com/mmcdole/gofeed"
)

const REDDIT = "reddit"

type RedditProvider struct {
	*Proxy
}

func NewRedditProvider(p *Proxy) Provider {
	return &RedditProvider{
		Proxy: p,
	}
}

func (p *RedditProvider) Convert(item *gofeed.Item, model *models.ItemModel) {
	model.Provider = REDDIT
	model.Thumbnail = p.getThumb(item)
	link := p.extractLink(item.Content)
	if isImage(link) {
		model.Type = models.TypeImage
	}
	if isVideo(link) || p.isVideoProvider(link) {
		model.Type = models.TypeVideo
	}
}

func (p *RedditProvider) Fetch(item *models.ItemModel) {

	item.Description = item.Content
	item.Link = p.extractLink(item.Content)

	//avoid recursion
	if p.Proxy.registers.GetRegisterByURL(item.Link).Name != REDDIT {
		p.Proxy.Fetch(item)
	}
}

func (p *RedditProvider) Download(item *models.ItemModel) {
	item.Link = p.extractLink(item.Content)
}

func (p *RedditProvider) getThumb(item *gofeed.Item) string {
	if item.Extensions != nil {
		return item.Extensions["media"]["thumbnail"][0].Attrs["url"]
	}
	return ""
}

func (p *RedditProvider) isVideoProvider(src string) bool {
	//TODO fix this
	return p.Proxy.registers.GetRegisterByURL(src).Name == YOUTUBE
}

func (p *RedditProvider) extractLink(content string) string {

	content = html.UnescapeString(content)
	pattern := `<a href\s*=\s*["\']?([^"\'\s>]+)["\']?>\[link\]`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return ""
	}

	found := re.FindAllStringSubmatch(content, 2)
	if len(found) > 0 {
		return found[0][1]
	}

	return ""
}
