package data

import (
	"krolus/models"
	"time"
)

// SubscriptionManager ...
type SubscriptionManager interface {
	Add(sub *models.SubscriptionModel) error
	Remove(ID string) error
	Get(ID string) (*models.SubscriptionModel, error)
	GetByURL(XURL string) (*models.SubscriptionModel, error)
	Update(sub *models.SubscriptionModel) error
	AllByIDs(IDs ...string) (models.SubscriptionCollection, error)
	ForEachOlderThan(time.Duration, func(*models.SubscriptionModel) error) error
}

// ItemManager ...
type ItemManager interface {
	Add(item *models.ItemModel) error
	AddInBatch(models.SubscriptionItemsMap) error
	AllPaginated(models.PaginatedRequest) (models.PaginatedItemCollection, error)
	Get(string) (*models.ItemModel, error)
	GetUpdate(string) (*models.ItemModel, error)
	UpdateFavorite(string) error
}

// Manager ...
type Manager struct {
	Subscription SubscriptionManager
	Item         ItemManager
}
