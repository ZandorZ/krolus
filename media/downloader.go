package media

import (
	"html"
	"krolus/models"
	"regexp"
	"time"

	readability "github.com/go-shiori/go-readability"
	"github.com/pkg/errors"
)

// Downloader ...
type Downloader struct {
}

// Download ...
func (d *Downloader) Download(item *models.ItemModel) (string, error) {

	var link string
	switch item.Provider {
	case "reddit":
		if url, ok := redditExtractLink(item.Description); ok {
			link = url
		}
	default:
		link = item.Link
	}

	article, err := readability.FromURL(link, 30*time.Second)
	if err != nil {
		return "", errors.Wrapf(err, "can't download: %s", link)
	}

	return article.Content, nil
}

func redditExtractLink(content string) (string, bool) {

	content = html.UnescapeString(content)

	re, err := regexp.Compile(`<a href\s*=\s*["\']?([^"\'\s>]+)["\']?>\[link\]`)

	if err != nil {
		panic(err)
	}

	found := re.FindAllStringSubmatch(content, 2)

	if len(found) > 0 {
		return found[0][1], true
	}

	return "", false
}
