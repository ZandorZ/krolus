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
	return s.Model(sub).Where("id = ?", sub.ID).Updates(&sub).Error
}

// Remove ...
func (s *SubscriptionManagerSqte) Remove(ID string) error {
	return nil
}

// Get ...
func (s *SubscriptionManagerSqte) Get(ID string) (*models.SubscriptionModel, error) {
	var sub models.SubscriptionModel
	err := s.DB.First(&sub, "id = ?", ID).Error
	return &sub, err
}

// GetByURL ...
func (s *SubscriptionManagerSqte) GetByURL(XURL string) (*models.SubscriptionModel, error) {
	var sub models.SubscriptionModel
	err := s.DB.First(&sub, "xurl = ?", XURL).Error
	return &sub, err
}

// AllByIDs ...
func (s *SubscriptionManagerSqte) AllByIDs(IDs ...string) (models.SubscriptionCollection, error) {
	var subs []models.SubscriptionModel
	err := s.DB.Find(&subs, IDs).Error
	return subs, err
}

// ForEachOlderThan ...
func (s *SubscriptionManagerSqte) ForEachOlderThan(since time.Duration, forEachFn func(*models.SubscriptionModel) error) error {
	rows, err := s.DB.Model(&models.SubscriptionModel{}).
		Order("last_updated DESC").
		Where("last_updated < ?", time.Now().Add(-since)).
		Rows()

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var sub models.SubscriptionModel
		if err := s.DB.ScanRows(rows, &sub); err != nil {
			return err
		} else {
			forEachFn(&sub)
		}
	}
	return nil
}
