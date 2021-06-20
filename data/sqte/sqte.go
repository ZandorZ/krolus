package sqte

import (
	"krolus/data"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dataManager *data.Manager

func NewManager(path string, debug bool) *data.Manager {

	if dataManager == nil {
		config := &gorm.Config{}
		if debug {
			config.Logger = logger.Default.LogMode(logger.Warn)
		}

		db, err := gorm.Open(sqlite.Open(path), config)
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
