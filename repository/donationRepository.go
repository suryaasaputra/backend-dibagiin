package repository

import (
	"dibagi/models"

	"gorm.io/gorm"
)

type IDonationRepository interface {
	Create(models.Donation) (models.CreateDonationResponse, error)
}

type DonationDb struct {
	db *gorm.DB
}

func NewDonationRepository(db *gorm.DB) *DonationDb {
	return &DonationDb{
		db: db,
	}
}

func (d DonationDb) Create(donation models.Donation) (models.CreateDonationResponse, error) {
	err := d.db.Create(&donation).Error
	if err != nil {
		return models.CreateDonationResponse{}, err
	}
	return models.CreateDonationResponse{
		ID:          donation.ID,
		UserID:      donation.UserID,
		Title:       donation.Title,
		Description: donation.Description,
		PhotoUrl:    donation.PhotoUrl,
		Status:      donation.Status,
		Location:    donation.Location,
		CreatedAt:   donation.CreatedAt,
	}, nil
}
