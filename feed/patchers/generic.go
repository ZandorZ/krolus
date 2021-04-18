package patchers

import (
	"krolus/models"
	"time"

	"github.com/mmcdole/gofeed"
)

// Patch ...
func GenericPatch(item *gofeed.Item) *models.ItemModel {

	img := ""
	if item.Image != nil {
		img = item.Image.URL
	}
	// else if len(item.Enclosures) > 0 && item.Enclosures[0].Type == "image/jpeg" {
	// 	img = item.Enclosures[0].URL
	// }

	d := item.PublishedParsed
	if d == nil {
		d = item.UpdatedParsed
	}

	var date time.Time
	if d != nil {
		date = d.Local()
	}

	return &models.ItemModel{
		Title:       item.Title,
		Link:        item.Link,
		Description: item.Description,
		Thumbnail:   img,
		Published:   date,
		New:         true,
		Provider:    "generic",
		Type:        "unknown",
	}
}
