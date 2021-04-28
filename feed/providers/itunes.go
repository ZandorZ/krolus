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

	newItem := &models.ItemModel{
		ID:          uuid.New().String(),
		Title:       item.Title,
		Description: item.Description,
		Content:     item.Content,
		New:         true,
		Thumbnail:   item.ITunesExt.Image,
		Published:   item.PublishedParsed.Local(),
		Provider:    "itunes",
	}

	p.patchLink(newItem, item)
	p.patchText(newItem, item)

	return newItem
}

func (p *ITunesProvider) Fetch(item *models.ItemModel) {

}

func (p *ITunesProvider) Download(item *models.ItemModel) error {
	return nil
}

func (p *ITunesProvider) patchLink(newItem *models.ItemModel, item *gofeed.Item) {

	if len(item.Enclosures) > 0 && item.Enclosures[0].Type == "audio/mpeg" {
		newItem.Link = item.Enclosures[0].URL
		newItem.Type = models.TypeAudio
	}

	if len(item.Enclosures) > 0 && item.Enclosures[0].Type == "video/mpeg" {
		newItem.Link = item.Enclosures[0].URL
		newItem.Type = models.TypeVideo
	}
}

func (p *ITunesProvider) patchText(newItem *models.ItemModel, item *gofeed.Item) {
	if newItem.Description == "" && newItem.Content == "" {
		newItem.Description = item.ITunesExt.Summary
	}
}
