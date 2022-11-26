package repository

import (
	"dibagi/models"

	"gorm.io/gorm"
)

type IDonationRequestRepository interface {
	GetAllByUserId(string) ([]models.GetDonationRequestResponse, error)
	GetAllByDonationId(string) ([]models.GetDonationRequestResponse, error)
	GetAllByDonatorId(string) ([]models.GetDonationRequestResponse, error)
	GetById(string) (models.GetDonationRequestResponse, error)
	GetByDonationId(string) (models.DonationRequest, error)
	Create(models.DonationRequest) (models.CreateDonationRequestResponse, error)
	Confirm(string) error
	Reject(string) error
}

type DonationRequestDb struct {
	db *gorm.DB
}

func NewDonationRequestRepository(db *gorm.DB) *DonationRequestDb {
	return &DonationRequestDb{
		db: db,
	}
}

func (d DonationRequestDb) Create(donationRequest models.DonationRequest) (models.CreateDonationRequestResponse, error) {
	donationModel := models.Donation{}
	donationId := donationRequest.DonationID

	err := d.db.Where("id=?", donationId).First(&donationModel).Error
	if err != nil {
		return models.CreateDonationRequestResponse{}, err
	}

	donationRequest.DonatorID = donationModel.UserID

	err = d.db.Create(&donationRequest).Error
	if err != nil {
		return models.CreateDonationRequestResponse{}, err
	}
	return models.CreateDonationRequestResponse{
		ID:         donationRequest.ID,
		UserID:     donationRequest.UserID,
		DonationID: donationRequest.DonationID,
		Message:    donationRequest.Message,
		Status:     donationRequest.Status,
		CreatedAt:  donationRequest.CreatedAt,
	}, nil
}

func (d DonationRequestDb) GetAllByUserId(userId string) ([]models.GetDonationRequestResponse, error) {
	donationRequests := []models.DonationRequest{}
	err := d.db.Preload("User").Preload("Donation").Preload("Donation.User").Where("user_id=?", userId).Find(&donationRequests).Error
	if err != nil {
		return nil, err
	}
	listDonationRequest := []models.GetDonationRequestResponse{}
	for _, v := range donationRequests {
		response := models.GetDonationRequestResponse{
			ID:         v.ID,
			UserID:     v.UserID,
			DonationID: v.DonationID,
			DonatorID:  v.DonatorID,
			Status:     v.Status,
			Message:    v.Message,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		}

		response.Donation.ID = v.Donation.ID
		response.Donation.Title = v.Donation.Title
		response.Donation.Description = v.Donation.Description
		response.Donation.PhotoUrl = v.Donation.PhotoUrl
		response.Donation.Location = v.Donation.Location
		response.Donation.CreatedAt = v.Donation.CreatedAt
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
		listDonationRequest = append(listDonationRequest, response)
	}
	return listDonationRequest, nil
}

func (d DonationRequestDb) GetAllByDonationId(donationId string) ([]models.GetDonationRequestResponse, error) {
	donationRequests := []models.DonationRequest{}
	err := d.db.Preload("User").Preload("Donation").Preload("Donation.User").Where("donation_id=?", donationId).Find(&donationRequests).Error
	if err != nil {
		return nil, err
	}
	listDonationRequest := []models.GetDonationRequestResponse{}
	for _, v := range donationRequests {
		response := models.GetDonationRequestResponse{
			ID:         v.ID,
			UserID:     v.UserID,
			DonationID: v.DonationID,
			DonatorID:  v.DonatorID,
			Status:     v.Status,
			Message:    v.Message,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		}

		response.Donation.ID = v.Donation.ID
		response.Donation.Title = v.Donation.Title
		response.Donation.Description = v.Donation.Description
		response.Donation.PhotoUrl = v.Donation.PhotoUrl
		response.Donation.Location = v.Donation.Location
		response.Donation.CreatedAt = v.Donation.CreatedAt
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
		listDonationRequest = append(listDonationRequest, response)
	}
	return listDonationRequest, nil
}
func (d DonationRequestDb) GetAllByDonatorId(donatorId string) ([]models.GetDonationRequestResponse, error) {
	donationRequests := []models.DonationRequest{}
	err := d.db.Preload("User").Preload("Donation").Preload("Donation.User").Where("donator_id=?", donatorId).Find(&donationRequests).Error
	if err != nil {
		return nil, err
	}
	listDonationRequest := []models.GetDonationRequestResponse{}
	for _, v := range donationRequests {
		response := models.GetDonationRequestResponse{
			ID:         v.ID,
			UserID:     v.UserID,
			DonationID: v.DonationID,
			DonatorID:  v.DonatorID,
			Status:     v.Status,
			Message:    v.Message,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		}

		response.Donation.ID = v.Donation.ID
		response.Donation.Title = v.Donation.Title
		response.Donation.Description = v.Donation.Description
		response.Donation.PhotoUrl = v.Donation.PhotoUrl
		response.Donation.Location = v.Donation.Location
		response.Donation.CreatedAt = v.Donation.CreatedAt
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

		listDonationRequest = append(listDonationRequest, response)
	}
	return listDonationRequest, nil
}

func (d DonationRequestDb) GetById(id string) (models.GetDonationRequestResponse, error) {
	donationRequest := models.DonationRequest{}
	err := d.db.Where("id=?", id).Preload("User").Preload("Donation").Preload("Donation.User").First(&donationRequest).Error
	if err != nil {
		return models.GetDonationRequestResponse{}, err
	}

	result := models.GetDonationRequestResponse{
		ID:         donationRequest.ID,
		UserID:     donationRequest.UserID,
		DonationID: donationRequest.DonationID,
		DonatorID:  donationRequest.DonatorID,
		Status:     donationRequest.Status,
		Message:    donationRequest.Message,
		CreatedAt:  donationRequest.CreatedAt,
		UpdatedAt:  donationRequest.UpdatedAt,
	}
	result.Donation.ID = donationRequest.Donation.ID
	result.Donation.Title = donationRequest.Donation.Title
	result.Donation.Description = donationRequest.Donation.Description
	result.Donation.PhotoUrl = donationRequest.Donation.PhotoUrl
	result.Donation.Location = donationRequest.Donation.Location
	result.Donation.CreatedAt = donationRequest.Donation.CreatedAt
	result.Donation.Donator.ID = donationRequest.Donation.User.ID
	result.Donation.Donator.UserName = donationRequest.Donation.User.UserName
	result.Donation.Donator.FullName = donationRequest.Donation.User.FullName
	result.Donation.Donator.PhoneNumber = donationRequest.Donation.User.PhoneNumber
	result.Donation.Donator.ProfilPhotoUrl = donationRequest.Donation.User.ProfilPhotoUrl

	result.User.ID = donationRequest.User.ID
	result.User.UserName = donationRequest.User.UserName
	result.User.FullName = donationRequest.User.FullName
	result.User.PhoneNumber = donationRequest.User.PhoneNumber
	result.User.ProfilPhotoUrl = donationRequest.User.ProfilPhotoUrl

	return result, nil
}

func (d DonationRequestDb) GetByDonationId(id string) (models.DonationRequest, error) {
	donationRequest := models.DonationRequest{}
	err := d.db.Where("donation_id=?", id).First(&donationRequest).Error
	if err != nil {
		return models.DonationRequest{}, err
	}
	return donationRequest, nil
}

func (d DonationRequestDb) Confirm(id string) error {
	donationRequest := models.DonationRequest{}

	err := d.db.Where("id=?", id).First(&donationRequest).Error
	if err != nil {
		return err
	}

	err = d.db.Model(&donationRequest).Where("id=?", id).Update("status", "confirmed").Error
	if err != nil {
		return err
	}

	donation := models.Donation{}
	donationID := donationRequest.DonationID

	err = d.db.Model(&donation).Where("id=?", donationID).Updates(map[string]interface{}{"status": "taken", "taker_id": donationRequest.UserID}).Error
	if err != nil {
		return err
	}

	donationHistory := models.DonationHistory{
		DonationID: donationID,
		UserID:     donationRequest.UserID,
	}

	err = d.db.Create(&donationHistory).Error
	if err != nil {
		return err
	}

	return nil
}
func (d DonationRequestDb) Reject(id string) error {
	donationRequest := models.DonationRequest{}

	err := d.db.Where("id=?", id).First(&donationRequest).Error
	if err != nil {
		return err
	}

	err = d.db.Model(&donationRequest).Where("id=?", id).Update("status", "rejected").Error
	if err != nil {
		return err
	}
	return nil
}
