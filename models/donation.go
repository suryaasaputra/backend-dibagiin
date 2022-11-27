package models

import (
	"mime/multipart"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/rs/xid"
	"gorm.io/gorm"
)

type Donation struct {
	ID              string            `json:"id" gorm:"primaryKey;type:varchar"`
	UserID          string            `json:"-" gorm:"not null;"`
	Title           string            `json:"title" form:"title" gorm:"not null;type:varchar" valid:"required~Title is required"`
	Description     string            `json:"description" form:"description" gorm:"not null;type:varchar" valid:"required~Description is required"`
	PhotoUrl        string            `json:"photo_url" form:"photo_url" gorm:"not null;type:varchar" valid:"required~Photo URL is required"`
	Location        string            `json:"location" form:"location" gorm:"not null;type:varchar" valid:"required~Location is required"`
	Status          string            `json:"status" gorm:"not null;type:varchar;default:available"`
	TakerID         *string           `json:"taker_id" gorm:"type:varchar;"`
	DonationRequest []DonationRequest `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User            *User             `json:"-"`
	Taker           *User             `json:"-"`
	CreatedAt       *time.Time        `json:"created_at"`
	UpdatedAt       *time.Time        `json:"updated_at"`
}

type CreateDonationRequest struct {
	Title         string                `json:"title" form:"title"`
	Description   string                `json:"description" form:"description"`
	DonationPhoto *multipart.FileHeader `json:"donation_photo" form:"donation_photo"`
	Location      string                `json:"location" form:"location"`
}
type EditDonationRequest struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Location    string `json:"location" form:"location"`
}

type CreateDonationResponse struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	Title       string     `json:"title" `
	Description string     `json:"description"`
	PhotoUrl    string     `json:"photo"`
	Location    string     `json:"location"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
}

type EditDonationResponse struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	Title       string     `json:"title" `
	Description string     `json:"description"`
	PhotoUrl    string     `json:"photo"`
	Location    string     `json:"location"`
	Status      string     `json:"status"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
type GetDonationsResponse struct {
	Donation
	Request []string `json:"requester_id"`
	Donator struct {
		ID             string `json:"id"`
		UserName       string `json:"user_name"`
		FullName       string `json:"full_name"`
		PhoneNumber    string `json:"phone_number"`
		ProfilPhotoUrl string `json:"profil_photo_url"`
	} `json:"donator"`
	Taker struct {
		ID             string `json:"id,omitempty"`
		UserName       string `json:"user_name,omitempty"`
		FullName       string `json:"full_name,omitempty"`
		PhoneNumber    string `json:"phone_number,omitempty"`
		ProfilPhotoUrl string `json:"profil_photo_url,omitempty"`
	} `json:"taker"`
}

func (d *Donation) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(d)
	if err != nil {
		return err
	}
	newId := xid.New().String()
	d.ID = newId
	return nil
}
