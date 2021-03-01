package patchers

import (
	"fmt"
	"krolus/models"

	"github.com/mmcdole/gofeed"
)

// Bitchute ...
type Bitchute struct {
}

// NewBitchutePatcher ....
func NewBitchutePatcher() Patcher {
	return &Bitchute{}
}

// Patch ...
func (l *Bitchute) Patch(item *gofeed.Item) *models.ItemModel {

	return &models.ItemModel{
		ID:          item.GUID,
		Title:       item.Title,
		Link:        fmt.Sprintf("https://www.bitchute.com/embed/%s/", item.GUID),
		Description: item.Description,
		Thumbnail:   item.Enclosures[0].URL,
		Published:   item.PublishedParsed.Local(),
		New:         true,
		Provider:    "bitchute",
		Type:        "video",
		Embed:       item.Link,
	}
}
