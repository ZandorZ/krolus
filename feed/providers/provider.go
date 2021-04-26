package providers

import (
	"krolus/models"

	"github.com/mmcdole/gofeed"
)

///////////////////////////////////////////////////////////////////
type ConvertFunc func(*models.SubscriptionModel, *gofeed.Feed) models.ItemCollection
type Converter interface {
	Convert(*models.SubscriptionModel, *gofeed.Feed) models.ItemCollection
}

//////////////////////////////////////////////////////////////////
type PatchFunc func(*models.ItemModel)
type Patcher interface {
	Patch(*models.ItemModel)
}

//////////////////////////////////////////////////////////////////
type FetcherFunc func(*Proxy, *models.ItemModel)
type Fetcher interface {
	Fetch(*Proxy, *models.ItemModel)
}

////////////////////////////////////////////////////////////////
type Downloader interface {
	Download(*models.ItemModel) error
}
