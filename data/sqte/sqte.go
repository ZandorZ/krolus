package sqte

import (
	"krolus/data"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dataManager *data.Manager

func NewManager(path string) *data.Manager {

	if dataManager == nil {
		db, err := gorm.Open(sqlite.Open(path), &gorm.Config{
			// SkipDefaultTransaction: true,
		})
		if err != nil {
			panic("failed to connect database")
		}
		dataManager = &data.Manager{
			Subscription: newSubscriptionManagerSqte(db),
			Item:         newItemManagerSqte(db),
		}
	}

	return dataManager
}
