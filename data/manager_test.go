package data_test

import (
	"krolus/data"
	"krolus/data/sqte"
	"krolus/models"
	"testing"

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
		title := "Sub test"
		err := manager.Subscription.Add(&models.SubscriptionModel{
			ID:          id,
			Title:       title,
			Description: "This is a test",
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

}
