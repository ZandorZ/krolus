package feed

import (
	"krolus/models"
	"sync"
)

// Observable ...
type Observable <-chan models.SubscriptionItemsMap

// SubInfos information about new items by sub
type SubInfos struct {
	lock  sync.RWMutex
	infos models.SubscriptionItemsMap
}

// NewSubInfos ...
func NewSubInfos() *SubInfos {
	return &SubInfos{
		infos: make(models.SubscriptionItemsMap),
	}
}

// Put ...
func (s *SubInfos) Put(sub *models.SubscriptionModel, items *models.ItemCollection) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.infos[sub] = items
}

// Len ...
func (s *SubInfos) Len() int {
	return len(s.infos)
}

// Reset ...
func (s *SubInfos) Reset() {
	s.infos = make(models.SubscriptionItemsMap)
}

// Infos ...
func (s *SubInfos) Infos() models.SubscriptionItemsMap {
	return s.infos
}

// Observer ...
type Observer interface {
	Add(string) Observable
	Publish(info models.SubscriptionItemsMap)
}

// MyObserver ...
type MyObserver struct {
	lock     sync.RWMutex
	channels map[string]chan models.SubscriptionItemsMap
}

// NewObserver ....
func NewObserver() Observer {
	return &MyObserver{
		channels: make(map[string]chan models.SubscriptionItemsMap),
	}
}

// Add observer
func (o *MyObserver) Add(key string) Observable {
	o.lock.Lock()
	defer o.lock.Unlock()
	if o.channels[key] == nil {
		o.channels[key] = make(chan models.SubscriptionItemsMap)
	}
	return o.channels[key]
}

// Publish to observables
func (o *MyObserver) Publish(info models.SubscriptionItemsMap) {
	o.lock.Lock()
	defer o.lock.Unlock()
	for _, obs := range o.channels {
		obs <- info
	}
}
