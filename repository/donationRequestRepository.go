package repository

import (
	"dibagi/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	donationHistory := models.DonationHistory{
		DonationID:        donationId,
		UserID:            donationModel.UserID,
		DonationRequestID: donationRequest.ID,
		Type:              "request",
		Message:           "Meminta barang yang anda donasikan",
	}

	err = d.db.Create(&donationHistory).Error
	if err != nil {
		return models.CreateDonationRequestResponse{}, err
	}

	return models.CreateDonationRequestResponse{
		ID:         donationRequest.ID,
		UserID:     donationRequest.UserID,
		DonationID: donationRequest.DonationID,
		DonatorID:  donationRequest.DonatorID,
		Message:    donationRequest.Message,
		Status:     donationRequest.Status,
		CreatedAt:  donationRequest.CreatedAt,
	}, nil
}

func (d DonationRequestDb) GetAllByUserId(userId string) ([]models.GetDonationRequestResponse, error) {
	donationRequests := []models.DonationRequest{}
	err := d.db.Preload("User").Preload("Donator").Preload("Donation").Where("user_id=?", userId).Order("created_at desc").Find(&donationRequests).Error
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
		response.Donation.UpdatedAt = v.Donation.UpdatedAt

		response.Donator.ID = v.Donator.ID
		response.Donator.UserName = v.Donator.UserName
		response.Donator.FullName = v.Donator.FullName
		response.Donator.PhoneNumber = v.Donator.PhoneNumber
		response.Donator.ProfilPhotoUrl = v.Donator.ProfilPhotoUrl

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
	err := d.db.Preload("User").Preload("Donator").Preload("Donation").Where("donation_id=?", donationId).Find(&donationRequests).Error
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
		response.Donation.UpdatedAt = v.Donation.UpdatedAt

		response.Donator.ID = v.Donator.ID
		response.Donator.UserName = v.Donator.UserName
		response.Donator.FullName = v.Donator.FullName
		response.Donator.PhoneNumber = v.Donator.PhoneNumber
		response.Donator.ProfilPhotoUrl = v.Donator.ProfilPhotoUrl

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
	err := d.db.Preload("User").Preload("Donator").Preload("Donation").Where("donator_id=?", donatorId).Order("created_at desc").Find(&donationRequests).Error
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
		response.Donation.UpdatedAt = v.Donation.UpdatedAt

		response.Donator.ID = v.Donator.ID
		response.Donator.UserName = v.Donator.UserName
		response.Donator.FullName = v.Donator.FullName
		response.Donator.PhoneNumber = v.Donator.PhoneNumber
		response.Donator.ProfilPhotoUrl = v.Donator.ProfilPhotoUrl

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
	err := d.db.Preload("User").Preload("Donator").Preload("Donation").Where("id=?", id).First(&donationRequest).Error
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
	result.Donation.UpdatedAt = donationRequest.Donation.UpdatedAt

	result.Donator.ID = donationRequest.Donator.ID
	result.Donator.UserName = donationRequest.Donator.UserName
	result.Donator.FullName = donationRequest.Donator.FullName
	result.Donator.PhoneNumber = donationRequest.Donator.PhoneNumber
	result.Donator.ProfilPhotoUrl = donationRequest.Donator.ProfilPhotoUrl

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

	err = d.db.Model(&donationRequest).Where("id=?", id).Update("status", "Dikonfirmasi").Error
	if err != nil {
		return err
	}

	err = d.db.Model(models.DonationRequest{}).Where("donation_id=?", donationRequest.DonationID).Not("user_id=?", donationRequest.UserID).Updates(models.DonationRequest{Status: "Ditolak"}).Error
	if err != nil {
		return err
	}

	listRejectedRequest := []models.DonationRequest{}
	err = d.db.Where("donation_id=?", donationRequest.DonationID).Where("status=?", "Ditolak").Find(&listRejectedRequest).Error
	if err != nil {
		return err
	}

	for _, v := range listRejectedRequest {
		donationHistory := models.DonationHistory{
			DonationID:        v.DonationID,
			UserID:            v.UserID,
			DonationRequestID: id,
			Type:              "rejectAll",
			Message:           "donasi telah diberikan ke pengguna lain",
		}
		err = d.db.Create(&donationHistory).Error
		if err != nil {
			return err
		}
	}

	donation := models.Donation{}
	donationID := donationRequest.DonationID

	err = d.db.Model(&donation).Where("id=?", donationID).Updates(map[string]interface{}{"status": "Sudah diambil", "taker_id": donationRequest.UserID}).Error
	if err != nil {
		return err
	}

	donationHistory := models.DonationHistory{
		DonationID:        donationID,
		UserID:            donationRequest.UserID,
		DonationRequestID: donationRequest.ID,
		Type:              "confirm",
		Message:           "Permintaan anda telah disetujui",
	}

	err = d.db.Create(&donationHistory).Error
	if err != nil {
		return err
	}

	return nil
}

func (d DonationRequestDb) Reject(id string) error {
	donationRequest := models.DonationRequest{}

	err := d.db.Model(&donationRequest).Where("id=?", id).Updates(map[string]interface{}{"status": "Ditolak"}).Error
	if err != nil {
		return err
	}

	err = d.db.Preload(clause.Associations).Where("id=?", id).First(&donationRequest).Error
	if err != nil {
		return err
	}

	donationHistory := models.DonationHistory{
		DonationID:        donationRequest.DonationID,
		UserID:            donationRequest.UserID,
		DonationRequestID: donationRequest.ID,
		Type:              "reject",
		Message:           "menolak permintaan anda",
	}
	err = d.db.Create(&donationHistory).Error
	if err != nil {
		return err
	}
	return nil
}
