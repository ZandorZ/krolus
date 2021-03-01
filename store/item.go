package store

import (
	"krolus/data"
	"krolus/models"
	"krolus/treex"

	"github.com/pkg/browser"
)

// ItemStore ...
type ItemStore struct {
	manager   *data.Manager
	treeState *treex.State
}

// NewItemStore ...
func NewItemStore(manager *data.Manager, treeState *treex.State) IItemStore {
	return &ItemStore{
		manager:   manager,
		treeState: treeState,
	}
}

// FetchItem ...
func (i *ItemStore) FetchItem(itemID string, updateNew bool) (models.ItemModel, error) {
	var item *models.ItemModel
	var err error

	if updateNew {
		if item, err = i.manager.Item.GetUpdate(itemID); err != nil {
			return *item, err
		}
		//TODO: fix this ///////////
		i.treeState.MapLeaves.Get(item.Subscription).NewItemsCount--
		if err := i.treeState.Save(); err != nil {
			return models.ItemModel{}, err
		}
		//////////////////////////////////////////////////
	} else {
		item, err = i.manager.Item.Get(itemID)
		return *item, err
	}

	return *item, err
}

// OpenItem ...
func (i *ItemStore) OpenItem(url string) error {
	return browser.OpenURL(url)

}
