package providers

import (
	"net/url"
	"time"

	"github.com/mmcdole/gofeed"
)

func getImage(item *gofeed.Item) string {
	img := ""
	if item.Image != nil {
		img = item.Image.URL
	} else if len(item.Enclosures) > 0 && item.Enclosures[0].Type == "image/jpeg" {
		img = item.Enclosures[0].URL
	}
	return img
}

func getDate(item *gofeed.Item) time.Time {
	d := item.PublishedParsed
	if d == nil {
		d = item.UpdatedParsed
	}
	if d == nil {
		d = &time.Time{}
	}
	return d.Local()
}

func getDomain(link string) string {
	u, _ := url.Parse(link)
	return u.Hostname()
}
