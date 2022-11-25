package repository

import (
	"dibagi/models"

	"gorm.io/gorm"
)

type IDonationHistoryRepository interface {
	GetAll() ([]models.DonationHistory, error)
}

type DonationHistoryDB struct {
	db *gorm.DB
}

func NewDonationHistoryRepository(db *gorm.DB) *DonationHistoryDB {
	return &DonationHistoryDB{
		db: db,
	}
}

func (d DonationHistoryDB) GetAll() ([]models.DonationHistory, error) {

	return []models.DonationHistory{}, nil
}
