package media

import (
	"time"

	readability "github.com/go-shiori/go-readability"
)

// Downloader ...
type Downloader struct {
}

// Download ...
func (d *Downloader) Download(url string) (string, error) {
	article, err := readability.FromURL(url, 30*time.Second)
	if err != nil {
		return "", err
	}

	//  article.Image

	return article.Content, nil
}
