package db

import (
	"fmt"
	"krolus/models"
	"time"

	"github.com/timshannon/badgerhold/v3"
)

// SubscriptionManagerBH ...
type SubscriptionManagerBH struct {
	*badgerhold.Store
}

// Add ...
func (s *SubscriptionManagerBH) Add(sub *models.SubscriptionModel) error {
	return s.Insert(sub.ID, sub)
}

// Update ...
func (s *SubscriptionManagerBH) Update(sub *models.SubscriptionModel) error {

	//TODO: rename SubscriptionName in Items[]

	return s.Store.UpdateMatching(&models.SubscriptionModel{}, badgerhold.Where("ID").Eq(sub.ID), func(record interface{}) error {
		update, ok := record.(*models.SubscriptionModel)
		if !ok {
			return fmt.Errorf("Record isn't the correct type!  Wanted models.SubscriptionModel, got %T", record)
		}
		update.AlertNewItems = sub.AlertNewItems
		update.Title = sub.Title
		update.Description = sub.Description

		return nil
	})
}

// Remove ...
func (s *SubscriptionManagerBH) Remove(ID string) error {
	return s.Delete(ID, &models.SubscriptionModel{})
}

// Get ...
func (s *SubscriptionManagerBH) Get(ID string) (*models.SubscriptionModel, error) {
	sub := &models.SubscriptionModel{}
	err := s.Store.Get(ID, sub)
	return sub, err
}

// GetByURL ...
func (s *SubscriptionManagerBH) GetByURL(XURL string) (*models.SubscriptionModel, error) {
	sub := &models.SubscriptionModel{}
	err := s.Store.FindOne(sub, badgerhold.Where("XURL").Eq(XURL))
	return sub, err
}

// AllByIDs ...
func (s *SubscriptionManagerBH) AllByIDs(IDs ...string) (models.SubscriptionCollection, error) {
	// TODO: foreach, dont load all into memory
	var result models.SubscriptionCollection
	err := s.Find(&result,
		badgerhold.
			Where("ID").
			In(stringsToGenerics(IDs...)...).
			SortBy("Title"))
	return result, err
}

// ForEachOlderThan ...
func (s *SubscriptionManagerBH) ForEachOlderThan(since time.Duration, forEachFn func(*models.SubscriptionModel) error) error {
	return s.Store.ForEach(badgerhold.
		Where("LastUpdate").
		Lt(time.Now().Add(-since)).
		SortBy("LastUpdate").
		Reverse(), forEachFn)
}
