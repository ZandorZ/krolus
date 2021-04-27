package providers

import (
	"krolus/models"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
)

type GenericProvider struct {
	*Proxy
}

func NewGenericProvider(p *Proxy) Provider {
	return &GenericProvider{
		Proxy: p,
	}
}

func (g *GenericProvider) Convert(item *gofeed.Item) *models.ItemModel {

	//special case
	if item.ITunesExt != nil {
		return g.registers.GetRegisterByKey("itunes").Provide(g.Proxy).Convert(item)
	}

	return &models.ItemModel{
		ID:          uuid.New().String(),
		Title:       item.Title,
		Link:        item.Link,
		Description: item.Description,
		Content:     item.Content,
		New:         true,
		Thumbnail:   getImage(item),
		Published:   getDate(item),
		Provider:    "generic",
		Type:        models.TypeUnknown, //TODO: verify link (image, video, pdf, mp3)
	}
}

func (g *GenericProvider) Fetch(item *models.ItemModel) {

}

func (g *GenericProvider) Download(item *models.ItemModel) error {
	return nil
}
