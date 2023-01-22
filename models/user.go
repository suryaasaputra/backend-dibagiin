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
	Lat              float64           ` json:"lat" form:"lat" valid:"required~Latitude is required"`
	Lng              float64           ` json:"lng" form:"lng" valid:"required~Longitude is required"`
	ProfilPhotoUrl   string            `gorm:"type:varchar;" json:"profil_photo_url" form:"profil_photo_url" `
	CreatedAt        *time.Time        `json:"created_at"`
	UpdatedAt        *time.Time        `json:"updated_at"`
	Donation         []Donation        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`
	TakenDonation    []Donation        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:TakerID"`
	DonationRequest  []DonationRequest `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`
	SubmittedRequest []DonationRequest `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:DonatorID"`
	Notification     []Notification    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserRegisterRequest struct {
	Email       string  `json:"email" form:"email" example:"user@mail.com"`
	UserName    string  `json:"user_name" form:"user_name" example:"user_name"`
	Password    string  `json:"password" form:"password" example:"pass1234"`
	FullName    string  `json:"full_name" form:"full_name" example:"Full Name" `
	Gender      string  `json:"gender" form:"gender" example:"Male"`
	PhoneNumber string  `json:"phone_number" form:"phone_number" example:"+62890123456" `
	Address     string  `json:"address" form:"address" example:"Jakarta"`
	Lat         float64 `json:"lat" form:"lng" example:"-6.20104" `
	Lng         float64 `json:"lng" form:"lng" example:"106.816666" `
	// ProfilPhoto *multipart.FileHeader `json:"profil_photo,omitempty" form:"profil_photo,omitempty" `
}
type EditUserRequest struct {
	Email       string  `json:"email" form:"email" example:"user@mail.com"`
	UserName    string  `json:"user_name" form:"user_name" example:"user_name"`
	FullName    string  `json:"full_name" form:"full_name" example:"Full Name" `
	Gender      string  `json:"gender" form:"gender" example:"Male"`
	PhoneNumber string  `json:"phone_number" form:"phone_number" example:"+62890123456" `
	Address     string  `json:"address" form:"address" example:"Jakarta"`
	Lat         float64 `json:"lat" form:"lng" example:"-6.20104" `
	Lng         float64 `json:"lng" form:"lng" example:"106.816666" `
	// ProfilPhoto *multipart.FileHeader `json:"profil_photo,omitempty" form:"profil_photo,omitempty" `
}
type UserLoginRequest struct {
	Email    string `json:"email" form:"email" example:"tes@mail.com"`
	Password string `json:"password" form:"password" example:"pass1234"`
}

type SetProfilePhotoRequest struct {
	ProfilPhoto *multipart.FileHeader `json:"profil_photo,omitempty" form:"profil_photo,omitempty" `
}
type CreateUserResponse struct {
	ID             string     `json:"id" `
	Email          string     `json:"email" `
	UserName       string     `json:"user_name" `
	FullName       string     `json:"full_name" `
	Gender         string     `json:"gender" `
	PhoneNumber    string     `json:"phone_number" `
	Address        string     `json:"address"`
	Lat            float64    `json:"lat" form:"lat" `
	Lng            float64    `json:"lng" form:"lng" `
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
	Lat            float64    `json:"lat" form:"lat" `
	Lng            float64    `json:"lng" form:"lng" `
	ProfilPhotoUrl string     `json:"profil_photo_url"`
	Updated_at     *time.Time `json:"updated_at"`
}
type LoginUserResponse struct {
	ID             string     `json:"id"`
	Email          string     `json:"email"`
	UserName       string     `json:"user_name"`
	FullName       string     `json:"full_name"`
	PhoneNumber    string     `json:"phone_number"`
	Address        string     `json:"address"`
	Lat            float64    `json:"lat" form:"lat" `
	Lng            float64    `json:"lng" form:"lng" `
	ProfilPhotoUrl string     `json:"profil_photo_url"`
	LoginTime      *time.Time `json:"login_time"`
	Token          string     `json:"token"`
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
		ID              string            `json:"id"`
		Title           string            `json:"title"`
		Description     string            `json:"description"`
		Weight          int               `json:"weight"`
		PhotoUrl        string            `json:"photo_url"`
		Location        string            `json:"location"`
		Status          string            `json:"status"`
		Request         []string          `json:"requester_id"`
		DonationRequest []DonationRequest `json:"-"`
		TakerID         *string           `json:"taker_id"`
		CreatedAt       *time.Time        `json:"created_at"`
		UpdatedAt       *time.Time        `json:"updated_at"`
		Donator         struct {
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
		if u.Gender == "wanita" {
			u.ProfilPhotoUrl = config.STORAGE_PATH + "/profil_photo/default-girl.png"
		} else {
			u.ProfilPhotoUrl = config.STORAGE_PATH + "/profil_photo/default-man.png"
		}
	}
	return nil
}
