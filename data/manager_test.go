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
	manager = sqte.NewManager("file:test.db?cache=shared&mode=memory")
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
		title := "Sub test"
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
				Title:       fmt.Sprintf("Syb Test %d", i),
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
		err := manager.Subscription.ForEachOlderThan(since, func(sm *models.SubscriptionModel) error {
			assert.IsType(t, &models.SubscriptionModel{}, sm)
			assert.True(t, sm.LastUpdate.Before(time.Now().Add(-since)))
			counter++
			return nil
		})
		assert.Nil(t, err)
		assert.Equal(t, 4, counter)
	})

}
