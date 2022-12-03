package repository

import (
	"dibagi/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IDonationRepository interface {
	Create(models.Donation) (models.CreateDonationResponse, error)
	GetAll() ([]models.GetDonationsResponse, error)
	GetById(string) (models.GetDonationsResponse, error)
	GetAllAvailable() ([]models.GetDonationsResponse, error)
	GetAllByLocation(string) ([]models.GetDonationsResponse, error)
	Edit(string, models.EditDonationRequest) (models.EditDonationResponse, error)
	Delete(string) error
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

func (d DonationDb) GetAll() ([]models.GetDonationsResponse, error) {
	donations := []models.Donation{}
	err := d.db.Preload(clause.Associations).Order("created_at desc").Find(&donations).Error
	if err != nil {
		return nil, err
	}

	donationList := []models.GetDonationsResponse{}

	for _, v := range donations {
		response := models.GetDonationsResponse{}
		response.Donation = v
		response.Donator.ID = v.User.ID
		response.Donator.UserName = v.User.UserName
		response.Donator.FullName = v.User.FullName
		response.Donator.PhoneNumber = v.User.PhoneNumber
		response.Donator.ProfilPhotoUrl = v.User.ProfilPhotoUrl
		for _, r := range v.DonationRequest {
			response.Request = append(response.Request, r.UserID)
		}

		if response.TakerID != nil {
			response.Taker.ID = v.User.ID
			response.Taker.UserName = v.Taker.UserName
			response.Taker.FullName = v.Taker.FullName
			response.Taker.PhoneNumber = v.Taker.PhoneNumber
			response.Taker.ProfilPhotoUrl = v.Taker.ProfilPhotoUrl
		}
		donationList = append(donationList, response)
	}

	return donationList, nil
}

func (d DonationDb) GetAllAvailable() ([]models.GetDonationsResponse, error) {
	donations := []models.Donation{}
	err := d.db.Where("status=?", "Tersedia").Preload("User").Find(&donations).Error
	if err != nil {
		return nil, err
	}

	donationList := []models.GetDonationsResponse{}

	for _, v := range donations {
		response := models.GetDonationsResponse{}
		response.Donation = v
		response.Donator.ID = v.User.ID
		response.Donator.UserName = v.User.UserName
		response.Donator.FullName = v.User.FullName
		response.Donator.ProfilPhotoUrl = v.User.ProfilPhotoUrl
		donationList = append(donationList, response)
	}
	return donationList, nil
}
func (d DonationDb) GetAllByLocation(location string) ([]models.GetDonationsResponse, error) {
	donations := []models.Donation{}
	err := d.db.Where("location LIKE", "%"+location+"%").Preload("User").Find(&donations).Error
	if err != nil {
		return nil, err
	}

	donationList := []models.GetDonationsResponse{}

	for _, v := range donations {
		response := models.GetDonationsResponse{}
		response.Donation = v
		response.Donator.ID = v.User.ID
		response.Donator.UserName = v.User.UserName
		response.Donator.FullName = v.User.FullName
		response.Donator.ProfilPhotoUrl = v.User.ProfilPhotoUrl
		donationList = append(donationList, response)
	}
	return donationList, nil
}

func (d DonationDb) GetById(id string) (models.GetDonationsResponse, error) {
	donation := models.Donation{}
	err := d.db.Where("id=?", id).Preload("User").Preload("Taker").First(&donation).Error
	if err != nil {
		return models.GetDonationsResponse{}, err
	}
	var donator = struct {
		ID             string `json:"id"`
		UserName       string `json:"user_name"`
		FullName       string `json:"full_name"`
		PhoneNumber    string `json:"phone_number"`
		ProfilPhotoUrl string `json:"profil_photo_url"`
	}{}
	donator.ID = donation.User.ID
	donator.UserName = donation.User.UserName
	donator.FullName = donation.User.FullName
	donator.PhoneNumber = donation.User.PhoneNumber
	donator.ProfilPhotoUrl = donation.User.ProfilPhotoUrl
	var taker = struct {
		ID             string `json:"id,omitempty"`
		UserName       string `json:"user_name,omitempty"`
		FullName       string `json:"full_name,omitempty"`
		PhoneNumber    string `json:"phone_number,omitempty"`
		ProfilPhotoUrl string `json:"profil_photo_url,omitempty"`
	}{}
	if donation.TakerID != nil {
		taker.ID = donation.Taker.ID
		taker.UserName = donation.Taker.UserName
		taker.FullName = donation.Taker.FullName
		taker.PhoneNumber = donation.Taker.PhoneNumber
		taker.ProfilPhotoUrl = donation.Taker.ProfilPhotoUrl
	}
	result := models.GetDonationsResponse{
		Donation: donation,
		Donator:  donator,
		Taker:    taker,
	}
	return result, nil
}

func (d DonationDb) Edit(id string, new_data models.EditDonationRequest) (models.EditDonationResponse, error) {
	donationModel := models.Donation{}
	err := d.db.Model(&donationModel).Clauses(clause.Returning{}).
		Where("id=?", id).Updates(models.Donation{
		Title:       new_data.Title,
		Description: new_data.Description,
		Location:    new_data.Location,
	}).Error

	if err != nil {
		return models.EditDonationResponse{}, err
	}
	response := models.EditDonationResponse{
		ID:          donationModel.ID,
		UserID:      donationModel.UserID,
		PhotoUrl:    donationModel.PhotoUrl,
		Status:      donationModel.Status,
		Title:       donationModel.Title,
		Description: donationModel.Description,
		Location:    donationModel.Location,
		UpdatedAt:   donationModel.UpdatedAt,
	}
	return response, nil
}

func (d DonationDb) Delete(id string) error {
	Donation := models.Donation{}
	err := d.db.Where("id=?", id).Delete(&Donation).Error
	if err != nil {
		return err
	}
	return nil
}
