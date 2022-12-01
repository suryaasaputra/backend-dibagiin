package repository

import (
	"dibagi/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IDonationHistoryRepository interface {
	GetAllByUserId(string) ([]models.GetDonationHistoryResponse, error)
	GetAllByDonationRequestId(string) ([]models.DonationHistory, error)
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
	err := d.db.Preload("User").Preload("Donation").Preload("Donation.User").Preload("DonationRequest").Preload("DonationRequest.User").Where("user_id=?", userId).Order("created_at desc").Find(&donationHistory).Error
	if err != nil {
		return nil, err
	}
	listHistory := []models.GetDonationHistoryResponse{}
	for _, v := range donationHistory {
		response := models.GetDonationHistoryResponse{
			ID:                v.ID,
			UserID:            v.UserID,
			DonationID:        v.DonationID,
			Type:              v.Type,
			Message:           v.Message,
			DonationRequestID: v.DonationRequestID,
			CreatedAt:         v.CreatedAt,
			UpdatedAt:         v.UpdatedAt,
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
		response.User.ProfilPhotoUrl = v.User.ProfilPhotoUrl

		response.DonationRequest.ID = v.DonationRequest.ID
		response.DonationRequest.UserID = v.DonationRequest.UserID
		response.DonationRequest.DonationID = v.DonationRequest.DonationID
		response.DonationRequest.DonatorID = v.DonationRequest.DonatorID
		response.DonationRequest.Status = v.DonationRequest.Status
		response.DonationRequest.Message = v.DonationRequest.Message
		response.DonationRequest.CreatedAt = v.DonationRequest.CreatedAt
		response.DonationRequest.UpdatedAt = v.DonationRequest.UpdatedAt

		response.DonationRequest.User.ID = v.DonationRequest.User.ID
		response.DonationRequest.User.FullName = v.DonationRequest.User.FullName
		response.DonationRequest.User.UserName = v.DonationRequest.User.UserName
		response.DonationRequest.User.PhoneNumber = v.DonationRequest.User.PhoneNumber
		response.DonationRequest.User.ProfilPhotoUrl = v.DonationRequest.User.ProfilPhotoUrl

		listHistory = append(listHistory, response)
	}
	return listHistory, nil
}
func (d DonationHistoryDB) GetAllByDonationRequestId(donationRequestId string) ([]models.DonationHistory, error) {
	donationHistory := []models.DonationHistory{}
	err := d.db.Preload(clause.Associations).Where("donation_request_id=?", donationRequestId).Find(&donationHistory).Error
	if err != nil {
		return nil, err
	}

	return donationHistory, nil
}
