package patchers

import (
	"krolus/models"

	"github.com/mmcdole/gofeed"
)

// Patcher ...
type Patcher interface {
	Patch(*gofeed.Item) *models.ItemModel
}

// PatchMap ...
var PatchMap = map[string]func() Patcher{
	"*":                   NewGenericPatcher,
	"www.youtube.com":     NewYoutubePatcher,
	"lbryfeed.melroy.org": NewLbryPatcher,
	"www.bitchute.com":    NewBitchutePatcher,
	"www.reddit.com":      NewRedditPatcher,
}
