package sqte

import (
	"krolus/models"

	"gorm.io/gorm"
)

type ItemManagerSqte struct {
	*baseSqte
}

func newItemManagerSqte(base *baseSqte) *ItemManagerSqte {
	// Migrate the schema
	base.AutoMigrate(&models.ItemModel{})

	return &ItemManagerSqte{base}
}

// Add saves an item
func (i *ItemManagerSqte) Add(item *models.ItemModel) error {
	return i.Create(item).Error
}

// AddInBatch ....
func (i *ItemManagerSqte) AddInBatch(subBatch models.SubscriptionItemsMap, _tx interface{}) error {

	tx, ok := _tx.(*gorm.DB)
	if !ok {
		tx = i.DB.Session(&gorm.Session{})
	}

	for sub, items := range subBatch {
		sliced := SplitItems(*items, 30)
		for _, slice := range sliced {
			if err := tx.CreateInBatches(slice, len(slice)).Error; err != nil {
				return err
			}
		} //date & link
		if err := tx.Model(sub).UpdateColumns(models.SubscriptionModel{Provider: sub.Provider, LastUpdate: sub.LastUpdate, LastItemLink: sub.LastItemLink}).Error; err != nil {
			return err
		}
	}
	return nil

}

// AllPaginated ...
func (i *ItemManagerSqte) AllPaginated(request models.PaginatedRequest) (models.PaginatedItemCollection, error) {

	var items models.ItemCollection
	itemsP := models.PaginatedItemCollection{
		Page:  request.Page,
		Items: items,
		Total: 0,
	}

	if request.LeafIDs == nil {
		return itemsP, nil
	}

	tx := i.tx
	if tx == nil {
		tx = i.DB.Session(&gorm.Session{})
	}

	query := tx.Model(&models.ItemModel{}).Select("id", "title", "published", "thumbnail", "subscription", "new", "favorite", "type").Preload("SubscriptionModel")

	if len(request.LeafIDs) > 0 {
		query.Where("Subscription IN (?)", request.LeafIDs)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return itemsP, err
	}

	if err := query.Limit(request.ItemsPerPage).
		Offset(request.ItemsPerPage * request.Page).
		Order("Published DESC").
		Find(&items).
		Error; err != nil {
		return itemsP, err
	}

	for i := range items {
		if items[i].SubscriptionModel != nil {
			items[i].SubscriptionName = items[i].SubscriptionModel.Title
		}
		items[i].SubscriptionModel = nil
	}

	itemsP.Items = items
	itemsP.Page = request.Page
	itemsP.Total = int(count)

	return itemsP, nil
}

// Get gets an item
func (i *ItemManagerSqte) Get(ID string) (*models.ItemModel, error) {
	var item models.ItemModel
	err := i.DB.First(&item, "id = ?", ID).Error
	return &item, err
}

// GetUpdate get and updates field New to false
func (i *ItemManagerSqte) GetUpdate(itemID string) (*models.ItemModel, error) {
	var item models.ItemModel

	tx := i.tx
	if tx == nil {
		tx = i.DB.Session(&gorm.Session{})
	}

	if err := tx.First(&item, "id = ?", itemID).Error; err != nil {
		return nil, err
	}

	if err := tx.Model(&item).Update("new", false).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

// UpdateFavorite updates the favorite field
func (i *ItemManagerSqte) UpdateFavorite(itemID string) error {
	tx := i.tx
	if tx == nil {
		tx = i.DB.Session(&gorm.Session{})
	}

	var item models.ItemModel
	if err := tx.First(&item, "id = ?", itemID).Error; err != nil {
		return err
	}

	return tx.Model(&item).Update("favorite", !item.Favorite).Error
}

func (i *ItemManagerSqte) All() (models.ItemCollection, error) {
	var items models.ItemCollection
	err := i.DB.Preload("SubscriptionModel").Find(&items).Error
	return items, err
}

// split items in chunks
func SplitItems(items models.ItemCollection, size int) []models.ItemCollection {
	slicedItems := []models.ItemCollection{}
	var j int
	for i := 0; i < len(items); i += size {
		j += size
		if j > len(items) {
			j = len(items)
		}
		slicedItems = append(slicedItems, items[i:j])
	}
	return slicedItems
}
