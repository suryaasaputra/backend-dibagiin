package repository

import (
	"dibagi/models"

	"gorm.io/gorm"
)

type IDonationRepository interface {
	Create(models.Donation) (models.CreateDonationResponse, error)
	GetDonations() ([]models.GetDonationsResponse, error)
	GetAvailableDonations() ([]models.GetDonationsResponse, error)
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

func (d DonationDb) GetDonations() ([]models.GetDonationsResponse, error) {
	donations := []models.Donation{}
	err := d.db.Find(&donations).Error
	if err != nil {
		return nil, err
	}

	donationList := []models.GetDonationsResponse{}

	for _, v := range donations {
		response := models.GetDonationsResponse{}
		user := models.User{}
		response.Donation = v
		donatorId := v.UserID
		err := d.db.Where("id=?", donatorId).First(&user).Error
		if err != nil {
			return nil, err
		}
		response.Donator = user.FullName
		donationList = append(donationList, response)

	}

	return donationList, nil
}
func (d DonationDb) GetAvailableDonations() ([]models.GetDonationsResponse, error) {
	donations := []models.Donation{}
	err := d.db.Where("status=?", "available").Find(&donations).Error
	if err != nil {
		return nil, err
	}

	donationList := []models.GetDonationsResponse{}

	for _, v := range donations {
		response := models.GetDonationsResponse{}
		user := models.User{}
		response.Donation = v
		donatorId := v.UserID
		err := d.db.Where("id=?", donatorId).First(&user).Error
		if err != nil {
			return nil, err
		}
		response.Donator = user.FullName
		donationList = append(donationList, response)

	}

	return donationList, nil
}
