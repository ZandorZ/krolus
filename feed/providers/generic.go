package providers

import (
	"krolus/models"

	"github.com/mmcdole/gofeed"
)

const GENERIC = "generic"

type GenericProvider struct {
	*Proxy
}

func NewGenericProvider(p *Proxy) Provider {
	return &GenericProvider{
		Proxy: p,
	}
}

func (g *GenericProvider) Convert(item *gofeed.Item, model *models.ItemModel) {
	model.Provider = GENERIC
	model.Type = models.TypeUnknown //TODO: verify link (image, video, pdf, mp3)
}

func (g *GenericProvider) Fetch(item *models.ItemModel) {

}

func (g *GenericProvider) Download(item *models.ItemModel) error {
	return nil
}
