package providers

import (
	"krolus/models"
	"time"

	"github.com/mmcdole/gofeed"
)

type Proxy struct {
	registers RegisterMap
}

func NewProxy() *Proxy {
	return &Proxy{
		registers: RegisterMap{
			"generic": &Register{
				Name:    "generic",
				Domains: []string{},
				Provide: NewGenericProvider,
			},
			"itunes": &Register{
				Name:    "itunes",
				Domains: []string{},
				Provide: NewItunesProvider,
			},
			"reddit": &Register{
				Name:    "reddit",
				Domains: []string{"reddit.com", "www.reddit.com"},
				Provide: NewRedditProvider,
			},
			"youtube": &Register{
				Name:    "youtube",
				Domains: []string{"youtu.be", "youtube.com", "www.youtube.com"},
				Provide: NewYoutubeProvider,
			},
		},
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
		firstTime = true
	}

	for _, item := range f.Items {

		newItem := *p.Convert(item)
		newItem.Subscription = sub.ID

		if firstTime {
			items = append(items, newItem)
			if latest.Published.Before(newItem.Published) {
				latest = newItem
			}
		} else if newItem.Link != sub.LastItemLink && newItem.Published.After(sub.LastUpdate) {
			items = append(items, newItem)
			if newItem.Published.After(latest.Published) {
				latest = newItem
			}
		}
	}

	if len(items) > 0 {
		sub.LastItemLink = latest.Link
		sub.LastUpdate = latest.Published
	}

	return items
}

func (p *Proxy) Convert(item *gofeed.Item) *models.ItemModel {
	return p.registers.GetRegisterByURL(item.Link).Provide(p).Convert(item)
}

func (p *Proxy) Fetch(item *models.ItemModel) {
	p.registers.GetRegisterByURL(item.Link).Provide(p).Fetch(item)
}

func (p *Proxy) Download(item *models.ItemModel) error {
	return p.registers.GetRegisterByURL(item.Link).Provide(p).Download(item)
}
