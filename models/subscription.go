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
	XURL          string
	Description   string
	URL           string
	LastUpdate    time.Time
	Status        uint
	AlertNewItems bool
}

// SubscriptionCollection ...
type SubscriptionCollection []SubscriptionModel

// SubscriptionItemsMap ...
type SubscriptionItemsMap map[*SubscriptionModel]*ItemCollection
