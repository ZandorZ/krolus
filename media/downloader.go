package media

import (
	"krolus/feed/providers"
	"krolus/models"
	"time"

	readability "github.com/go-shiori/go-readability"
	"github.com/pkg/errors"
)

// Downloader ...
type Downloader struct {
}

// Download ...
func (d *Downloader) Download(item *models.ItemModel) (string, error) {

	providers.NewProxy("..").Fetch(item)

	article, err := readability.FromURL(item.Link, 30*time.Second)
	if err != nil {
		return "", errors.Wrapf(err, "can't download: %s", item.Link)
	}

	return article.Content, nil
}
