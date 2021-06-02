package providers

import (
	"krolus/models"
	"time"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
)

type Proxy struct {
	registers RegisterMap
}

func NewProxy() *Proxy {

	registers := make(RegisterMap)
	registers.AddRegister(GENERIC, NewGenericProvider)
	registers.AddRegister(ITUNES, NewItunesProvider)
	registers.AddRegister(REDDIT, NewRedditProvider, "reddit.com", "www.reddit.com")
	registers.AddRegister(YOUTUBE, NewYoutubeProvider, "youtu.be", "youtube.com", "www.youtube.com")

	return &Proxy{
		registers: registers,
	}
}

// GetNewItems returns collection of new items from sub
func (p *Proxy) GetNewItems(sub *models.SubscriptionModel, f *gofeed.Feed) models.ItemCollection {

	items := models.ItemCollection{}
	latest := models.ItemModel{}
	firstTime := false

	//first time
	if time.Time.IsZero(sub.LastUpdate) {
		sub.Provider = p.registers.GetRegisterByURL(sub.XURL).Name
		if f.ITunesExt != nil {
			sub.Provider = ITUNES
		}
		firstTime = true
	}

	for _, item := range f.Items {

		newItem := p.newItemFrom(item)
		newItem.SubscriptionModel = sub
		p.Convert(item, newItem)
		newItem.Subscription = sub.ID
		newItem.SubscriptionModel = nil //TODO: weird

		if firstTime {
			items = append(items, *newItem)
			if latest.Published.Before(newItem.Published) {
				latest = *newItem
			}
		} else if newItem.Link != sub.LastItemLink && newItem.Published.After(sub.LastUpdate) {
			items = append(items, *newItem)
			if newItem.Published.After(latest.Published) {
				latest = *newItem
			}
		}
	}

	if len(items) > 0 {
		sub.LastItemLink = latest.Link
		sub.LastUpdate = latest.Published
	}

	return items
}

func (p *Proxy) newItemFrom(item *gofeed.Item) *models.ItemModel {
	return &models.ItemModel{
		ID:          uuid.New().String(),
		Title:       item.Title,
		Link:        item.Link,
		Description: item.Description,
		Content:     item.Content,
		Thumbnail:   getImage(item),
		Published:   getDate(item),
		New:         true,
		Type:        models.TypeUnknown,
	}
}

func (p *Proxy) Convert(item *gofeed.Item, model *models.ItemModel) {
	//special case
	if item.ITunesExt != nil || model.SubscriptionModel.Provider == ITUNES {
		p.registers.GetRegisterByKey(ITUNES).Provide(p).Convert(item, model)
		return
	}
	p.registers.GetRegisterByURL(item.Link).Provide(p).Convert(item, model)
}

func (p *Proxy) Fetch(item *models.ItemModel) {
	p.registers.GetRegisterByURL(item.Link).Provide(p).Fetch(item)
}

func (p *Proxy) Download(item *models.ItemModel) {
	p.registers.GetRegisterByURL(item.Link).Provide(p).Download(item)
}
