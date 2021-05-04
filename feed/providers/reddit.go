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
	if isImage(p.extractLink(item.Content)) {
		model.Type = models.TypeImage
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

func (p *RedditProvider) extractLink(content string) string {

	content = html.UnescapeString(content)
	// pattern := `<a href="(((https?\:\/\/)|(www\.))(\S+))">\[link\]</a>`
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

// func extractImage(content string) (string, bool) {

// 	content = html.UnescapeString(content)

// 	re, err := regexp.Compile(`<a href="((?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+(png|jpg|jpeg|gif|svg))">\[link\]</a>`)

// 	if err != nil {
// 		panic(err)
// 	}

// 	found := re.FindAllStringSubmatch(content, 2)

// 	if len(found) > 0 {
// 		return fmt.Sprintf("<img src='%s'>", found[0][1]), true
// 	}

// 	return "", false

// }
