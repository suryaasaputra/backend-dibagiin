package models

import (
	"dibagi/config"
	"dibagi/helpers"
	"mime/multipart"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/rs/xid"
	"gorm.io/gorm"
)

type User struct {
	ID               string            `gorm:"primaryKey;type:varchar" json:"id"`
	Email            string            `gorm:"not null;uniqueIndex;type:varchar" json:"email" form:"email" valid:"required~Email is required,email~Invalid email address"`
	UserName         string            `gorm:"not null;uniqueIndex;type:varchar" json:"user_name" form:"user_name" valid:"required~Username is required"`
	Password         string            `gorm:"not null;type:varchar" json:"password" form:"password" valid:"required~Password is required,minstringlength(8)"`
	FullName         string            `gorm:"not null;type:varchar" json:"full_name" form:"full_name" valid:"required~Full Name is required"`
	Gender           string            `gorm:"not null;type:varchar" json:"gender" form:"gender" valid:"required~Gender is required"`
	PhoneNumber      string            `gorm:"not null;type:varchar" json:"phone_number" form:"phone_number" valid:"required~Phone number is required"`
	Address          string            `gorm:"not null;type:varchar" json:"address" form:"address" valid:"required~Address is required"`
	ProfilPhotoUrl   string            `gorm:"type:varchar;" json:"profil_photo_url" form:"profil_photo_url" `
	CreatedAt        *time.Time        `json:"created_at"`
	UpdatedAt        *time.Time        `json:"updated_at"`
	Donation         []Donation        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`
	TakenDonation    []Donation        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:TakerID"`
	DonationRequest  []DonationRequest `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`
	SubmittedRequest []DonationRequest `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:DonatorID"`
	DonationHistory  []DonationHistory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserRegisterRequest struct {
	Email       string `json:"email" form:"email"`
	UserName    string `json:"user_name" form:"user_name"`
	Password    string `json:"password" form:"password" `
	FullName    string `json:"full_name" form:"full_name" `
	Gender      string `json:"gender" form:"gender" `
	PhoneNumber string `json:"phone_number" form:"phone_number" `
	Address     string `json:"address" form:"address" `
	// ProfilPhoto *multipart.FileHeader `json:"profil_photo,omitempty" form:"profil_photo,omitempty" `
}
type EditUserRequest struct {
	Email       string `json:"email" form:"email"`
	UserName    string `json:"user_name" form:"user_name"`
	FullName    string `json:"full_name" form:"full_name" `
	Gender      string `json:"gender" form:"gender" `
	PhoneNumber string `json:"phone_number" form:"phone_number" `
	Address     string `json:"address" form:"address" `
	// ProfilPhoto *multipart.FileHeader `json:"profil_photo,omitempty" form:"profil_photo,omitempty" `
}
type UserLoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type SetProfilePhotoRequest struct {
	ProfilPhoto *multipart.FileHeader `json:"profil_photo,omitempty" form:"profil_photo,omitempty" `
}
type CreateUserResponse struct {
	ID             string     `json:"id"`
	Email          string     `json:"email"`
	UserName       string     `json:"user_name"`
	FullName       string     `json:"full_name"`
	Gender         string     `json:"gender"`
	PhoneNumber    string     `json:"phone_number"`
	Address        string     `json:"address"`
	ProfilPhotoUrl string     `json:"profil_photo_url"`
	CreatedAt      *time.Time `json:"created_at"`
}
type EditUserResponse struct {
	ID             string     `json:"id"`
	Email          string     `json:"email"`
	UserName       string     `json:"user_name"`
	FullName       string     `json:"full_name"`
	Gender         string     `json:"gender"`
	PhoneNumber    string     `json:"phone_number"`
	Address        string     `json:"address"`
	ProfilPhotoUrl string     `json:"profil_photo_url"`
	Updated_at     *time.Time `json:"updated_at"`
}

type GetUserResponse struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	UserName       string `json:"user_name"`
	FullName       string `json:"full_name"`
	Gender         string `json:"gender"`
	PhoneNumber    string `json:"phone_number"`
	Address        string `json:"address"`
	ProfilPhotoUrl string `json:"profil_photo_url"`
	Donation       []struct {
		ID          string     `json:"id"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		PhotoUrl    string     `json:"photo_url"`
		Location    string     `json:"location"`
		Status      string     `json:"status"`
		TakerID     string     `json:"taker_id"`
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

	Created_at *time.Time `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
}

// Hooks model user
func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	hashedPassword, err := helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	newID := xid.New().String()
	u.ID = newID
	userName := strings.ReplaceAll(u.UserName, " ", "")
	lowerCase := strings.ToLower(userName)
	u.Email = strings.ToLower(u.Email)
	u.UserName = lowerCase
	u.Password = hashedPassword
	if u.ProfilPhotoUrl == "" {
		u.ProfilPhotoUrl = config.STORAGE_PATH + "/profil_photo/default.png"
	}
	return nil
}
