package providers

import (
	"krolus/models"

	"github.com/mmcdole/gofeed"
)

const ITUNES = "itunes"

type ITunesProvider struct {
	*Proxy
}

func NewItunesProvider(p *Proxy) Provider {
	return &ITunesProvider{
		Proxy: p,
	}
}

func (p *ITunesProvider) Convert(item *gofeed.Item, model *models.ItemModel) {
	model.Provider = ITUNES
	model.Thumbnail = item.ITunesExt.Image
	model.Published = item.PublishedParsed.Local()

	p.patchLink(item, model)
	p.patchText(item, model)
}

func (p *ITunesProvider) Fetch(item *models.ItemModel) {

}

func (p *ITunesProvider) Download(item *models.ItemModel) error {
	return nil
}

func (p *ITunesProvider) patchLink(item *gofeed.Item, newItem *models.ItemModel) {

	if len(item.Enclosures) > 0 && item.Enclosures[0].Type == "audio/mpeg" {
		newItem.Link = item.Enclosures[0].URL
		newItem.Type = models.TypeAudio
	}

	if len(item.Enclosures) > 0 && item.Enclosures[0].Type == "video/mpeg" {
		newItem.Link = item.Enclosures[0].URL
		newItem.Type = models.TypeVideo
	}
}

func (p *ITunesProvider) patchText(item *gofeed.Item, newItem *models.ItemModel) {
	if newItem.Description == "" && newItem.Content == "" {
		newItem.Description = item.ITunesExt.Summary
	}
}
