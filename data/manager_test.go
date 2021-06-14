package data_test

import (
	"fmt"
	"krolus/data"
	"krolus/data/sqte"
	"krolus/models"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var manager *data.Manager

func init() {
	manager = sqte.NewManager("file:test.db?cache=shared&mode=memory", true)
	// manager = bh.NewManager("./")
}

func Test_Subscription(t *testing.T) {

	t.Run("Add sub", func(t *testing.T) {

		id := uuid.New().String()
		title := "First test"
		err := manager.Subscription.Add(&models.SubscriptionModel{
			ID:          id,
			Title:       title,
			Description: "This is the first test",
			LastUpdate:  time.Now().Add(-time.Hour * 1),
		})
		assert.Nil(t, err)

		// get sub
		sub, err := manager.Subscription.Get(id)
		assert.Nil(t, err)

		// sub exist
		if assert.NotNil(t, sub) {
			// same title
			assert.Equal(t, title, sub.Title)
		}

	})

	t.Run("Update sub", func(t *testing.T) {

		id := uuid.New().String()
		sub := &models.SubscriptionModel{
			ID:          id,
			Title:       "Sub test",
			Description: "This is a test",
			LastUpdate:  time.Now().Add(-time.Hour * 2),
		}
		err := manager.Subscription.Add(sub)
		assert.Nil(t, err)

		sub.Title = "Sub test2"
		manager.Subscription.Update(sub)

		// get sub
		sub2, err := manager.Subscription.Get(id)
		assert.Nil(t, err)

		// sub exist
		if assert.NotNil(t, sub2) {
			// same title
			assert.Equal(t, sub2.Title, sub.Title)
		}

	})

	t.Run("Get sub by URL", func(t *testing.T) {

		id := uuid.New().String()
		title := "Sub test 3"
		url := "https://www.example.com"
		err := manager.Subscription.Add(&models.SubscriptionModel{
			ID:          id,
			Title:       title,
			Description: "This is a test",
			XURL:        url,
			LastUpdate:  time.Now().Add(-time.Hour * 3),
		})
		assert.Nil(t, err)

		// get sub
		sub, err := manager.Subscription.GetByURL(url)
		assert.Nil(t, err)

		// sub exist
		if assert.NotNil(t, sub) {
			// same title
			assert.Equal(t, title, sub.Title)
		}
	})

	t.Run("List all by IDs", func(t *testing.T) {

		ids := []string{uuid.New().String(), uuid.New().String(), uuid.New().String()}

		for i, id := range ids {
			err := manager.Subscription.Add(&models.SubscriptionModel{
				ID:          id,
				Title:       fmt.Sprintf("Sub Test %d", i),
				Description: fmt.Sprintf("This is a %d test", i),
				XURL:        fmt.Sprintf("http://www.example%d.com", i),
				LastUpdate:  time.Now().Add(-time.Hour * time.Duration(4+i)),
			})
			assert.Nil(t, err)
		}

		// find subs by id
		subs, err := manager.Subscription.AllByIDs(ids[0], ids[2])
		assert.Nil(t, err)

		assert.IsType(t, subs, models.SubscriptionCollection{})
		assert.Len(t, subs, 2)

	})

	t.Run("Foreach func", func(t *testing.T) {

		since := time.Hour * 3
		counter := 0
		err := manager.Subscription.ForEachOlderThan(since, func(sm *models.SubscriptionModel, _tx interface{}) error {
			assert.IsType(t, &models.SubscriptionModel{}, sm)
			assert.True(t, sm.LastUpdate.Before(time.Now().Add(-since)))
			counter++
			return nil
		})
		assert.Nil(t, err)
		assert.Equal(t, 4, counter)
	})

}

func Test_Item(t *testing.T) {

	t.Run("Add item", func(t *testing.T) {

		id := uuid.New().String()
		title := "First item test"
		err := manager.Item.Add(&models.ItemModel{
			Title:       title,
			Description: "This is the first test",
			Published:   time.Now().Add(-time.Hour * 1),
			New:         true,
		})
		assert.Nil(t, err)

		// get sub
		sub, err := manager.Item.Get(id)
		assert.Nil(t, err)

		// sub exist
		if assert.NotNil(t, sub) {
			// same title
			assert.Equal(t, title, sub.Title)
		}

	})

	t.Run("Add item in batch", func(t *testing.T) {

		itemsMap := models.SubscriptionItemsMap{}
		//get subs
		since := time.Hour * 3
		err := manager.Subscription.ForEachOlderThan(since, func(sm *models.SubscriptionModel, _tx interface{}) error {
			assert.IsType(t, &models.SubscriptionModel{}, sm)

			ids := []string{uuid.New().String(), uuid.New().String(), uuid.New().String()}
			var items models.ItemCollection
			for i, id := range ids {
				items = append(items, models.ItemModel{
					ID:                id,
					Title:             fmt.Sprintf("Item Test %d", i),
					Description:       fmt.Sprintf("This is a %d test", i),
					Link:              fmt.Sprintf("http://www.example%d.com", i),
					Published:         time.Now().Add(-time.Hour * time.Duration(2+i)),
					SubscriptionModel: sm,
					New:               true,
				})
			}
			sm.LastUpdate = items[len(items)-1].Published
			itemsMap[sm] = &items
			return nil
		})
		assert.Nil(t, err)

		err = manager.Item.AddInBatch(itemsMap, nil)
		assert.Nil(t, err)

	})

	t.Run("Query item pagination (without Sub IDs)", func(t *testing.T) {

		req := models.PaginatedRequest{
			Page:         0,
			ItemsPerPage: 3,
		}

		items, err := manager.Item.AllPaginated(req)
		assert.Nil(t, err)
		assert.Len(t, items.Items, 3)
		assert.Equal(t, 13, items.Total)

	})

	t.Run("Query item pagination (with Sub IDs)", func(t *testing.T) {

		//get subs
		since := time.Hour * 4
		var ids []string
		err := manager.Subscription.ForEachOlderThan(since, func(sm *models.SubscriptionModel, _tx interface{}) error {
			assert.IsType(t, &models.SubscriptionModel{}, sm)
			ids = append(ids, sm.ID)
			return nil
		})
		assert.Nil(t, err)

		req := models.PaginatedRequest{
			Page:         0,
			ItemsPerPage: 5,
			LeafIDs:      ids[0:2],
		}

		items, err := manager.Item.AllPaginated(req)
		assert.Nil(t, err)
		assert.Len(t, items.Items, 5)
		assert.Equal(t, 6, items.Total)

		req.Page = 1
		items, err = manager.Item.AllPaginated(req)
		assert.Nil(t, err)
		assert.Len(t, items.Items, 1)
		assert.Equal(t, 6, items.Total)

	})

	t.Run("GetUpated item", func(t *testing.T) {
		items, err := manager.Item.All()
		assert.Nil(t, err)

		item, err := manager.Item.GetUpdate(items[0].ID)
		assert.Nil(t, err)
		assert.False(t, item.New)

	})

	t.Run("Update favorite item", func(t *testing.T) {
		items, err := manager.Item.All()
		assert.Nil(t, err)

		err = manager.Item.UpdateFavorite(items[0].ID)
		assert.Nil(t, err)

		item, err := manager.Item.Get(items[0].ID)
		assert.Nil(t, err)
		assert.True(t, item.Favorite)

		err = manager.Item.UpdateFavorite(items[0].ID)
		assert.Nil(t, err)

		item, err = manager.Item.Get(items[0].ID)
		assert.Nil(t, err)
		assert.False(t, item.Favorite)

	})

	t.Run("Test sliced items", func(t *testing.T) {
		items, err := manager.Item.All()
		assert.Nil(t, err)

		sliced := sqte.SplitItems(items, 5)

		assert.Len(t, sliced, 3)

	})

}
