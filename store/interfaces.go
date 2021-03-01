package store

import (
	"krolus/models"
	"krolus/models/media"
)

// IFeedStore ...
type IFeedStore interface {
	LoadMoreItems(request map[string]interface{}) (models.PaginatedItemCollection, error)
	LoadSub(string) (models.SubscriptionModel, error)
	GetSub(string) (models.SubscriptionModel, error)
}

// ITreeStore interface
type ITreeStore interface {
	UnLoadAll() error
	LoadNode(ID string) error
	UnLoadNode(ID string) error
	MoveLeaf(leafID string, nodeID string) error
	MoveNode(nodeID string, toID string) error
	AddNode(node map[string]interface{}, toID string) error
	AddSubscription(sub map[string]interface{}, toID string) error
	RemoveSubscription(ID string) error
	EditSubscription(sub map[string]interface{}) error
	EditNode(node map[string]interface{}) error
}

// IMediaStore ...
type IMediaStore interface {
	FetchItem(string) (media.Media, error)
}

// IItemStore ...
type IItemStore interface {
	FetchItem(string, bool) (models.ItemModel, error)
	OpenItem(string) error
}
