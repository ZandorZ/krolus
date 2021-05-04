package models

import (
	"time"
)

// Status list
const (
	ACTIVE   = iota
	SUSPECT  = iota
	INACTIVE = iota
	DISABLED = iota
	MUTED    = iota
)

// SubscriptionModel ...
type SubscriptionModel struct {
	ID            string
	Title         string
	XURL          string `gorm:"column:xurl"`
	Description   string
	URL           string
	LastUpdate    time.Time `gorm:"column:last_updated"`
	LastItemLink  string    `gorm:"column:last_item"`
	Status        uint
	AlertNewItems bool
	Provider      string
}

// TableName overrides the table name
func (SubscriptionModel) TableName() string {
	return "subscriptions"
}

// SubscriptionCollection ...
type SubscriptionCollection []SubscriptionModel

// SubscriptionItemsMap ...
type SubscriptionItemsMap map[*SubscriptionModel]*ItemCollection
