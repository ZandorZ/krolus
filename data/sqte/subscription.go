package sqte

import (
	"krolus/models"
	"time"

	"gorm.io/gorm"
)

// SubscriptionManagerSqte ...
type SubscriptionManagerSqte struct {
	*gorm.DB
}

func newSubscriptionManagerSqte(db *gorm.DB) *SubscriptionManagerSqte {
	// Migrate the schema
	db.AutoMigrate(&models.SubscriptionModel{})

	return &SubscriptionManagerSqte{db}
}

// Add ...
func (s *SubscriptionManagerSqte) Add(sub *models.SubscriptionModel) error {
	return s.Create(sub).Error
}

// Update ...
func (s *SubscriptionManagerSqte) Update(sub *models.SubscriptionModel) error {
	return s.Model(sub).Where("ID = ?", sub.ID).Updates(&sub).Error
}

// Remove ...
func (s *SubscriptionManagerSqte) Remove(ID string) error {
	return nil
}

// Get ...
func (s *SubscriptionManagerSqte) Get(ID string) (*models.SubscriptionModel, error) {
	var sub models.SubscriptionModel
	err := s.First(&sub, "ID = ?", ID).Error
	return &sub, err
}

// GetByURL ...
func (s *SubscriptionManagerSqte) GetByURL(XURL string) (*models.SubscriptionModel, error) {
	return nil, nil
}

// AllByIDs ...
func (s *SubscriptionManagerSqte) AllByIDs(IDs ...string) (models.SubscriptionCollection, error) {
	return nil, nil
}

// ForEachOlderThan ...
func (s *SubscriptionManagerSqte) ForEachOlderThan(since time.Duration, forEachFn func(*models.SubscriptionModel) error) error {
	return nil
}
