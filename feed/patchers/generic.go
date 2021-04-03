package patchers

import (
	"krolus/models"
	"time"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
)

// Generic ...
type Generic struct {
}

// NewGenericPatcher ....
func NewGenericPatcher() Patcher {
	return &Generic{}
}

// Patch ...
func (g *Generic) Patch(item *gofeed.Item) *models.ItemModel {

	img := ""
	if item.Image != nil {
		img = item.Image.URL
	} else if len(item.Enclosures) > 0 && item.Enclosures[0].Type == "image/jpeg" {
		img = item.Enclosures[0].URL
	}

	d := item.PublishedParsed
	if d == nil {
		d = item.UpdatedParsed
	}

	var date time.Time
	if d != nil {
		date = d.Local()
	}

	return &models.ItemModel{
		ID:          uuid.New().String(),
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
