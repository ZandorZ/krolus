package sqte

import (
	"krolus/data"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dataManager *data.Manager

func NewManager(path string) *data.Manager {

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if dataManager == nil {
		dataManager = &data.Manager{
			Subscription: newSubscriptionManagerSqte(db),
			Item:         nil,
		}
	}

	return dataManager
}
