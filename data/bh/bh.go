package bh

import (
	"krolus/data"

	"github.com/timshannon/badgerhold/v3"
)

var (
	bhold       *badgerhold.Store
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
		bhold, err = badgerhold.Open(options)
		if err != nil {
			panic(err)
		}
		dataManager = &data.Manager{
			Subscription: &SubscriptionManagerBH{bhold},
			Item:         &ItemManagerBH{bhold},
		}
	}

	return dataManager
}

// ClearDB ...
func ClearDB() {
	bhold.Badger().DropAll()
}

// CloseDB ...
func CloseDB() {
	bhold.Badger().Close()
}

func stringsToGenerics(pars ...string) []interface{} {
	var out []interface{}
	for _, i := range pars {
		out = append(out, i)
	}
	return out
}
