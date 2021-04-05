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
	Status        uint
	AlertNewItems bool
}

// TableName overrides the table name
func (SubscriptionModel) TableName() string {
	return "subscriptions"
}

// SubscriptionCollection ...
type SubscriptionCollection []SubscriptionModel

// SubscriptionItemsMap ...
type SubscriptionItemsMap map[*SubscriptionModel]*ItemCollection
