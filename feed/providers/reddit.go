package providers

import (
	"html"
	"krolus/models"
	"regexp"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
)

type RedditProvider struct {
	*Proxy
}

func NewRedditProvider(p *Proxy) Provider {
	return &RedditProvider{
		Proxy: p,
	}
}

func (p *RedditProvider) Convert(item *gofeed.Item) *models.ItemModel {

	return &models.ItemModel{
		ID:          uuid.New().String(),
		Title:       item.Title,
		Link:        item.Link,
		Description: item.Description,
		Content:     item.Content,
		New:         true,
		Thumbnail:   p.getThumbReddit(item),
		Published:   item.PublishedParsed.Local(),
		Provider:    "reddit",
		Type:        "unknown",
	}
}

func (p *RedditProvider) Fetch(item *models.ItemModel) {
	item.Description = item.Content
	item.Link = p.extractLinkReddit(item.Content)

	//avoid recursion
	if p.Proxy.registers.GetRegisterByURL(item.Link).Name != "reddit" {
		p.Proxy.Fetch(item)
	}

}

func (p *RedditProvider) Download(item *models.ItemModel) error {
	return nil
}

func (p *RedditProvider) getThumbReddit(item *gofeed.Item) string {
	if item.Extensions != nil {
		return item.Extensions["media"]["thumbnail"][0].Attrs["url"]
	}
	return ""
}

func (p *RedditProvider) extractLinkReddit(content string) string {

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
