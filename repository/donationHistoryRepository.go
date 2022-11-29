package repository

import (
	"dibagi/models"

	"gorm.io/gorm"
)

type IDonationHistoryRepository interface {
	GetAllByUserId(string) ([]models.GetDonationHistoryResponse, error)
}

type DonationHistoryDB struct {
	db *gorm.DB
}

func NewDonationHistoryRepository(db *gorm.DB) *DonationHistoryDB {
	return &DonationHistoryDB{
		db: db,
	}
}

func (d DonationHistoryDB) GetAllByUserId(userId string) ([]models.GetDonationHistoryResponse, error) {
	donationHistory := []models.DonationHistory{}
	err := d.db.Preload("User").Preload("Donation").Preload("Donation.User").Where("user_id=?", userId).Order("created_at desc").Find(&donationHistory).Error
	if err != nil {
		return nil, err
	}
	listHistory := []models.GetDonationHistoryResponse{}
	for _, v := range donationHistory {
		response := models.GetDonationHistoryResponse{
			ID:         v.ID,
			UserID:     v.UserID,
			DonationID: v.DonationID,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		}

		response.Donation.ID = v.Donation.ID
		response.Donation.Title = v.Donation.Title
		response.Donation.Description = v.Donation.Description
		response.Donation.PhotoUrl = v.Donation.PhotoUrl
		response.Donation.Location = v.Donation.Location
		response.Donation.CreatedAt = v.Donation.CreatedAt
		response.Donation.UpdatedAt = v.Donation.UpdatedAt

		response.Donation.Donator.ID = v.Donation.User.ID
		response.Donation.Donator.UserName = v.Donation.User.UserName
		response.Donation.Donator.FullName = v.Donation.User.FullName
		response.Donation.Donator.PhoneNumber = v.Donation.User.PhoneNumber
		response.Donation.Donator.ProfilPhotoUrl = v.Donation.User.ProfilPhotoUrl

		response.User.ID = v.User.ID
		response.User.UserName = v.User.UserName
		response.User.FullName = v.User.FullName
		response.User.PhoneNumber = v.User.PhoneNumber

		listHistory = append(listHistory, response)
	}
	return listHistory, nil
}
