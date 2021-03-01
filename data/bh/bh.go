package db

import (
	"krolus/data"

	"github.com/timshannon/badgerhold/v3"
)

var (
	bh          *badgerhold.Store
	dataManager *data.Manager
)

// NewManager ...
func NewManager(path string) *data.Manager {

	if dataManager == nil {
		options := badgerhold.DefaultOptions
		options.Dir = path + "/_badger"
		options.ValueDir = path + "/_badger"
		options.SyncWrites = true

		var err error
		bh, err = badgerhold.Open(options)
		if err != nil {
			panic(err)
		}
		dataManager = &data.Manager{
			Subscription: &SubscriptionManagerBH{bh},
			Item:         &ItemManagerBH{bh},
		}
	}

	return dataManager
}

// ClearDB ...
func ClearDB() {
	bh.Badger().DropAll()
}

// CloseDB ...
func CloseDB() {
	bh.Badger().Close()
}

func stringsToGenerics(pars ...string) []interface{} {
	var out []interface{}
	for _, i := range pars {
		out = append(out, i)
	}
	return out
}
