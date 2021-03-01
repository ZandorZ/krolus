package models

import (
	"sync"

	"github.com/google/uuid"
)

// Leaf ...
type Leaf struct {
	lock          sync.RWMutex
	ID            string `json:"id"`
	ParentID      string `json:"parent"`
	Label         string `json:"label"`
	Description   string `json:"description"`
	Color         string `json:"color"`
	Icon          string `json:"icon"`
	ItemsCount    int    `json:"items_count"`
	NewItemsCount int    `json:"new_items_count"`
}

// NewLeaf ...
func NewLeaf(label, descripition string) *Leaf {
	return &Leaf{
		ID:          uuid.New().String(),
		Label:       label,
		Description: descripition,
	}
}
