package providers

import (
	"krolus/models"

	"github.com/mmcdole/gofeed"
)

//////////////////////////////////////////////////////////////////
type Converter interface {
	Convert(*gofeed.Item, *models.ItemModel)
}

//////////////////////////////////////////////////////////////////
type Fetcher interface {
	Fetch(*models.ItemModel)
}

////////////////////////////////////////////////////////////////
type Downloader interface {
	Download(*models.ItemModel)
}

type Provider interface {
	Converter
	Fetcher
	Downloader
}
