package providers

import (
	"krolus/models"
	"net/url"

	"github.com/mmcdole/gofeed"
)

var ConvertersMap map[string]ConvertFunc
var PatchersMap map[string]PatchFunc
var FetchersMap map[string]FetcherFunc

func init() {
	ConvertersMap = map[string]ConvertFunc{
		"*":               genericConverter,
		"www.reddit.com":  redditConverter,
		"www.youtube.com": yotubeConverter,
	}
	PatchersMap = map[string]PatchFunc{
		"*":               genericPatcher,
		"www.reddit.com":  redditPatcher,
		"www.youtube.com": youtubePatcher,
	}
	FetchersMap = map[string]FetcherFunc{
		"*":               genericFetcher,
		"www.youtube.com": youtubeFetcher,
	}
}

type Proxy struct {
	Patcher
	Fetcher
	ConvertFunc
}

func NewProxy(link string) *Proxy {
	return &Proxy{
		ConvertFunc: getConverter(getDomain(link)),
	}
}

func (p *Proxy) Convert(sub *models.SubscriptionModel, f *gofeed.Feed) models.ItemCollection {
	items := p.ConvertFunc(sub, f)
	for _, i := range items {
		p.Patch(&i)
	}
	return items
}

func (p *Proxy) Patch(item *models.ItemModel) {
	getPatcher(getDomain(item.Link))(item)
}

func (p *Proxy) Fetch(item *models.ItemModel) {
	getFetcher(getDomain(item.Link))(item)
}

func getConverter(domain string) ConvertFunc {
	conv := ConvertersMap["*"]
	if c, ok := ConvertersMap[domain]; ok {
		conv = c
	}
	return conv
}

func getPatcher(domain string) PatchFunc {
	patch := PatchersMap["*"]
	if p, ok := PatchersMap[domain]; ok {
		patch = p
	}
	return patch
}

func getFetcher(domain string) FetcherFunc {
	fetcher := FetchersMap["*"]
	if f, ok := FetchersMap[domain]; ok {
		fetcher = f
	}
	return fetcher
}

func getDomain(link string) string {
	u, _ := url.Parse(link)
	return u.Hostname()
}
