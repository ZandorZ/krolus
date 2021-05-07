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
		base := &baseSqte{DB: db}
		if err != nil {
			panic("failed to connect database")
		}
		dataManager = &data.Manager{
			Subscription: newSubscriptionManagerSqte(base),
			Item:         newItemManagerSqte(base),
		}
	}

	return dataManager
}

type baseSqte struct {
	*gorm.DB
	tx *gorm.DB
}

func (b *baseSqte) getTx() *gorm.DB {
	tx := b.tx
	if tx == nil {
		tx = b.DB
	}
	return tx
}
