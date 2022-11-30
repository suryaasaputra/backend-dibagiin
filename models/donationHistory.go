package models

import (
	"time"

	"github.com/rs/xid"
	"gorm.io/gorm"
)

type DonationHistory struct {
	ID                string           `json:"id" gorm:"primaryKey;type:varchar"`
	UserID            string           `json:"user_id" gorm:"not null;"`
	DonationID        string           `json:"donation_id" gorm:"not null;"`
	DonationRequestID string           `json:"donation_request_id" gorm:"not null;"`
	Status            bool             `json:"status" gorm:"not null;"`
	User              *User            `json:"-"`
	Donation          *Donation        `json:"-"`
	DonationRequest   *DonationRequest `json:"-"`
	CreatedAt         *time.Time       `json:"created_at"`
	UpdatedAt         *time.Time       `json:"updated_at"`
}

type GetDonationHistoryResponse struct {
	ID                string     `json:"id" `
	UserID            string     `json:"user_id"`
	DonationID        string     `json:"donation_id"`
	DonationRequestID string     `json:"donation_request_id"`
	Status            bool       `json:"status"`
	CreatedAt         *time.Time `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at"`
	Donation          struct {
		ID          string     `json:"id"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		PhotoUrl    string     `json:"photo_url"`
		Location    string     `json:"location"`
		CreatedAt   *time.Time `json:"created_at"`
		UpdatedAt   *time.Time `json:"updated_at"`
		Donator     struct {
			ID             string `json:"id"`
			UserName       string `json:"user_name"`
			FullName       string `json:"full_name"`
			PhoneNumber    string `json:"phone_number"`
			ProfilPhotoUrl string `json:"profil_photo_url"`
		} `json:"donator"`
	} `json:"donation"`
	User struct {
		ID             string `json:"id"`
		UserName       string `json:"user_name"`
		FullName       string `json:"full_name"`
		PhoneNumber    string `json:"phone_number"`
		ProfilPhotoUrl string `json:"profil_photo_url"`
	} `json:"user"`
}

func (d *DonationHistory) BeforeCreate(tx *gorm.DB) error {
	newId := xid.New().String()
	d.ID = newId
	return nil
}
