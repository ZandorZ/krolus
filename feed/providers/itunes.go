package providers

import (
	"krolus/models"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
)

type ITunesProvider struct {
	*Proxy
}

func NewItunesProvider(p *Proxy) Provider {
	return &ITunesProvider{
		Proxy: p,
	}
}

func (p *ITunesProvider) Convert(item *gofeed.Item) *models.ItemModel {

	return &models.ItemModel{
		ID:          uuid.New().String(),
		Title:       item.Title,
		Link:        p.itunesGetLink(item),
		Description: item.Description,
		Content:     item.Content,
		New:         true,
		Thumbnail:   item.ITunesExt.Image,
		Published:   item.PublishedParsed.Local(),
		Provider:    "itunes",
		Type:        models.TypeAudio,
	}
}

func (p *ITunesProvider) Fetch(item *models.ItemModel) {

}

func (p *ITunesProvider) Download(item *models.ItemModel) error {
	return nil
}

func (p *ITunesProvider) itunesGetLink(item *gofeed.Item) string {
	link := ""
	if len(item.Enclosures) > 0 && item.Enclosures[0].Type == "audio/mpeg" {
		link = item.Enclosures[0].URL
	}
	return link
}
