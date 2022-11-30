package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/rs/xid"
	"gorm.io/gorm"
)

type DonationRequest struct {
	ID         string     `json:"id" gorm:"primaryKey;type:varchar"`
	UserID     string     `json:"user_id" gorm:"not null;"`
	DonationID string     `json:"donation_id" gorm:"not null;"`
	DonatorID  string     `json:"donator_id" gorm:"not null;"`
	Message    string     `json:"message" gorm:"not null;" valid:"required~Message is required" `
	Status     string     `json:"status" gorm:"not null;type:varchar;default:Belum dikonfirmasi"`
	User       *User      `json:"-"`
	Donator    *User      `json:"-"`
	Donation   *Donation  `json:"-"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
type CreateDonationRequestRequest struct {
	Message string `json:"message"`
}
type CreateDonationRequestResponse struct {
	ID         string     `json:"id" `
	UserID     string     `json:"user_id"`
	DonationID string     `json:"donation_id"`
	DonatorID  string     `json:"donator_id"`
	Message    string     `json:"message"`
	Status     string     `json:"status"`
	CreatedAt  *time.Time `json:"created_at"`
}
type GetDonationRequestResponse struct {
	ID         string     `json:"id" `
	UserID     string     `json:"user_id"`
	DonationID string     `json:"donation_id"`
	DonatorID  string     `json:"donator_id"`
	Status     string     `json:"status"`
	Message    string     `json:"message"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	Donator    struct {
		ID             string `json:"id"`
		UserName       string `json:"user_name"`
		FullName       string `json:"full_name"`
		PhoneNumber    string `json:"phone_number"`
		ProfilPhotoUrl string `json:"profil_photo_url"`
	} `json:"donator"`
	Donation struct {
		ID          string     `json:"id"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		PhotoUrl    string     `json:"photo_url"`
		Location    string     `json:"location"`
		CreatedAt   *time.Time `json:"created_at"`
		UpdatedAt   *time.Time `json:"updated_at"`
	} `json:"donation"`
	User struct {
		ID             string `json:"id"`
		UserName       string `json:"user_name"`
		FullName       string `json:"full_name"`
		PhoneNumber    string `json:"phone_number"`
		ProfilPhotoUrl string `json:"profil_photo_url"`
	} `json:"user"`
}

func (d *DonationRequest) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(d)
	if err != nil {
		return err
	}
	newId := xid.New().String()
	d.ID = newId
	return nil
}
