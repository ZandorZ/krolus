package providers

import (
	"fmt"
	"net/url"
	"regexp"
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

func isImage(src string) bool {
	return isType(src, "png|jpg|jpeg|gif|svg|webp")
}

func isAudio(src string) bool {
	return isType(src, "mp3")
}

func isType(src string, extensions string) bool {
	re, err := regexp.Compile(fmt.Sprintf(`(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+(%s)`, extensions))
	if err != nil {
		return false
	}
	return re.MatchString(src)
}
