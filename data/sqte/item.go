package sqte

import (
	"krolus/models"

	"gorm.io/gorm"
)

type ItemManagerSqte struct {
	*gorm.DB
}

func newItemManagerSqte(db *gorm.DB) *ItemManagerSqte {
	// Migrate the schema
	db.AutoMigrate(&models.ItemModel{})

	return &ItemManagerSqte{db}
}

// Add saves an item
func (i *ItemManagerSqte) Add(item *models.ItemModel) error {
	return i.Create(item).Error
}

// AddInBatch ....
func (i *ItemManagerSqte) AddInBatch(subBatch models.SubscriptionItemsMap) error {

	for sub, items := range subBatch {

		sliced := SplitItems(*items, 30)
		for _, slice := range sliced {
			if err := i.DB.CreateInBatches(slice, len(slice)).Error; err != nil {
				return err
			}
		}
		if err := i.DB.Model(sub).UpdateColumn("LastUpdate", sub.LastUpdate).Error; err != nil {
			return err
		}
	}
	return nil

}

// AllPaginated ...
func (i *ItemManagerSqte) AllPaginated(request models.PaginatedRequest) (models.PaginatedItemCollection, error) {

	itemsP := models.PaginatedItemCollection{}
	var items models.ItemCollection

	return itemsP, i.DB.Transaction(func(tx *gorm.DB) error {

		query := tx.Model(&models.ItemModel{})

		if len(request.LeafIDs) > 0 {
			query = query.Where("subscription IN (?)", request.LeafIDs)
		}

		var count int64
		if err := query.Count(&count).Error; err != nil {
			return err
		}

		if err := query.Limit(request.ItemsPerPage).Offset(request.ItemsPerPage * request.Page).
			Find(&items).Error; err != nil {
			return err
		}

		itemsP.Items = items
		itemsP.Page = request.Page
		itemsP.Total = int(count)

		return nil

	})
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
	err := i.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&item, "id = ?", itemID).Error; err != nil {
			return err
		}
		return tx.Model(&item).Update("new", false).Error
	})
	return &item, err
}

// UpdateFavorite updates the favorite field
func (i *ItemManagerSqte) UpdateFavorite(itemID string) error {
	return i.DB.Transaction(func(tx *gorm.DB) error {
		var item models.ItemModel
		if err := tx.First(&item, "id = ?", itemID).Error; err != nil {
			return err
		}
		return tx.Model(&item).Update("favorite", !item.Favorite).Error
	})
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
