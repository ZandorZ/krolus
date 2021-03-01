package patchers

import (
	"fmt"
	"krolus/models"
	"regexp"

	"github.com/mmcdole/gofeed"
)

// Lbry ...
type Lbry struct {
}

// NewLbryPatcher ....
func NewLbryPatcher() Patcher {
	return &Lbry{}
}

func (l *Lbry) extractName(url string) (string, error) {

	re, err := regexp.Compile(`https:\/\/lbry.tv\/@([\w\.-]+):(\w)\/(.*):`)

	if err != nil {
		return "", err
	}

	found := re.FindAllStringSubmatch(url, 2)

	if len(found) > 0 {
		return found[0][3], nil
	}

	return "", fmt.Errorf("Patter not found")

}

// Patch ...
func (l *Lbry) Patch(item *gofeed.Item) *models.ItemModel {

	embed := ""
	name, err := l.extractName(item.Link)
	if err == nil {
		embed = fmt.Sprintf("https://lbry.tv/$/embed/%s/%s", name, item.GUID)
	}

	return &models.ItemModel{
		ID:          item.GUID,
		Title:       item.Title,
		Link:        item.Link,
		Description: item.Description,
		Thumbnail:   item.Enclosures[0].URL,
		Published:   item.PublishedParsed.Local(),
		New:         true,
		Provider:    "lbry",
		Type:        "video",
		Embed:       embed,
	}
}
