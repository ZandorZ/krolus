package db

import (
	"fmt"
	"krolus/models"

	"github.com/timshannon/badgerhold/v3"
)

// ItemManagerBH ...
type ItemManagerBH struct {
	*badgerhold.Store
}

// Add saves an item
func (i *ItemManagerBH) Add(item *models.ItemModel) error {
	return i.Insert(item.ID, item)
}

// AddInBatch ....
func (i *ItemManagerBH) AddInBatch(subBatch models.SubscriptionItemsMap) error {

	tx := i.Badger().NewTransaction(true)
	defer tx.Discard()

	for sub, items := range subBatch {
		for _, item := range *items {
			i.TxInsert(tx, item.ID, item)
		}
		// Update sub
		i.TxUpdate(tx, sub.ID, sub)
	}

	return tx.Commit()
}

// AllPaginated ...
func (i *ItemManagerBH) AllPaginated(request models.PaginatedRequest) (models.PaginatedItemCollection, error) {

	tx := i.Badger().NewTransaction(false)
	defer tx.Discard()

	subIDs := stringsToGenerics(request.LeafIDs...)
	var items models.ItemCollection
	var paginate models.PaginatedItemCollection

	q := badgerhold.Where("Subscription").
		In(subIDs...).
		Limit(request.ItemsPerPage).
		Skip(request.Page * request.ItemsPerPage).
		SortBy("Published").
		Reverse()

	if err := i.TxForEach(tx, q, func(item *models.ItemModel) error {
		// not all fields
		items = append(items, models.ItemModel{
			ID:               item.ID,
			Title:            item.Title,
			SubscriptionName: item.SubscriptionName,
			Subscription:     item.Subscription,
			Published:        item.Published,
			Thumbnail:        item.Thumbnail,
			Type:             item.Type,
			New:              item.New,
			Favorite:         item.Favorite,
		})
		return nil
	}); err != nil {
		return paginate, err
	}

	total, err := i.TxCount(tx, &models.ItemModel{},
		badgerhold.Where("Subscription").In(subIDs...))
	if err != nil {
		return paginate, err
	}

	paginate.Total = total
	paginate.Page = request.Page
	paginate.Items = items

	return paginate, nil
}

// Get gets an item
func (i *ItemManagerBH) Get(ID string) (*models.ItemModel, error) {
	item := &models.ItemModel{}
	err := i.Store.Get(ID, item)
	return item, err
}

// GetUpdate get and updates field New to false
func (i *ItemManagerBH) GetUpdate(itemID string) (*models.ItemModel, error) {

	tx := i.Store.Badger().NewTransaction(true)
	defer tx.Discard()

	item := &models.ItemModel{}
	if err := i.Store.TxGet(tx, itemID, item); err != nil {
		return nil, err
	}

	item.New = false
	if err := i.Store.TxUpsert(tx, item.ID, item); err != nil {
		return nil, err
	}

	return item, tx.Commit()
}

// UpdateFavorite updates the favorite field
func (i *ItemManagerBH) UpdateFavorite(itemID string) error {

	return i.Store.UpdateMatching(&models.ItemModel{}, badgerhold.Where("ID").Eq(itemID), func(record interface{}) error {

		update, ok := record.(*models.ItemModel) // record will always be a pointer
		if !ok {
			return fmt.Errorf("Record isn't the correct type!  Wanted *models.ItemModel, got %T", record)
		}
		update.Favorite = !update.Favorite

		return nil
	})
}
