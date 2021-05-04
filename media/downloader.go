package media

import (
	"krolus/feed"
	"krolus/models"

	readability "github.com/go-shiori/go-readability"
	"github.com/pkg/errors"
)

// Downloader ...
type Downloader struct {
	feed.Requester
}

func NewDownloader(req feed.Requester) *Downloader {
	return &Downloader{
		Requester: req,
	}
}

// Download ...
func (d *Downloader) Download(item *models.ItemModel) (string, error) {

	res, err := d.Request(item.Link)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	// TODO: use FromReader, custom client
	article, err := readability.FromReader(res.Body, item.Link)
	if err != nil {
		return "", errors.Wrapf(err, "can't download: %s", item.Link)
	}

	return article.Content, nil
}
