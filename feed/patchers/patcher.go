package patchers

import (
	"krolus/models"
	"net/url"

	"github.com/mmcdole/gofeed"
)

// Patcher ...
type Patcher func(*gofeed.Item) *models.ItemModel

// PatchMap ...
var PatchMap = map[string]Patcher{
	"*": GenericPatch,
}

// "www.youtube.com":  NewYoutubePatcher,
// 	"odysee.com":       NewLbryPatcher,
// 	"www.bitchute.com": NewBitchutePatcher,
// 	"www.reddit.com":   NewRedditPatcher,

func GetPatcher(link string) Patcher {

	u, _ := url.Parse(link)
	patcher := PatchMap["*"]
	if p, ok := PatchMap[u.Hostname()]; ok {
		patcher = p
	}
	return patcher
}
