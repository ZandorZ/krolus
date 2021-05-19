package models

import (
	"time"
)

// types
const (
	TypeText    = "text"
	TypeVideo   = "video"
	TypeAudio   = "audio"
	TypeImage   = "image"
	TypeUnknown = "unknown"
)

// ItemModel ...
type ItemModel struct {
	ID                string
	Title             string
	Link              string `gorm:"primaryKey"`
	Description       string
	Content           string
	Published         time.Time
	Subscription      string `gorm:"primaryKey"`
	SubscriptionName  string
	SubscriptionModel *SubscriptionModel `gorm:"foreignKey:Subscription" json:"-"`
	Provider          string
	Thumbnail         string
	Type              string
	Saved             bool
	New               bool
	Favorite          bool
}

// ItemCollection collections of items
type ItemCollection []ItemModel

// PaginatedItemCollection ...
type PaginatedItemCollection struct {
	Page  int
	Total int
	Items ItemCollection
}

// PaginatedRequest ...
type PaginatedRequest struct {
	Page         int
	ItemsPerPage int
	LeafIDs      []string
	NodeID       string
	Filter       *FilterRequest
}

// FilterRequest ...
type FilterRequest struct {
	New      *bool
	Favorite *bool
	Type     *[]string
}
