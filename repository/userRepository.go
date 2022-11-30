package repository

import (
	"dibagi/models"
	"errors"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUserRepository interface {
	Create(models.User) (models.CreateUserResponse, error)
	Edit(string, models.User) (models.EditUserResponse, error)
	GetByEmail(string) models.User
	GetByUserName(string) models.GetUserResponse
	GetById(string) models.GetUserResponse
	SetProfilePhoto(string, string) (models.EditUserResponse, error)
	Delete(string) error
}

type UserDb struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserDb {
	return &UserDb{
		db: db,
	}
}

func (u UserDb) Create(user models.User) (models.CreateUserResponse, error) {
	err := u.db.Create(&user).Error

	if err != nil {
		pgErr := err.(*pgconn.PgError)

		if pgErr.Code == pgerrcode.UniqueViolation {
			return models.CreateUserResponse{}, errors.New(pgErr.Detail)
		}
		return models.CreateUserResponse{}, err
	}

	return models.CreateUserResponse{
		ID:             user.ID,
		UserName:       user.UserName,
		Email:          user.Email,
		FullName:       user.FullName,
		Gender:         user.Gender,
		PhoneNumber:    user.PhoneNumber,
		Address:        user.Address,
		ProfilPhotoUrl: user.ProfilPhotoUrl,
		CreatedAt:      user.CreatedAt,
	}, nil
}

func (u UserDb) GetByEmail(email string) models.User {
	User := models.User{}

	u.db.Where("email =? ", email).First(&User)

	return User
}

func (u UserDb) GetByUserName(userName string) models.GetUserResponse {
	user := models.User{}

	u.db.Preload("Donation").Preload("Donation.User").Where("user_name =? ", userName).First(&user)
	response := models.GetUserResponse{
		ID:             user.ID,
		UserName:       user.UserName,
		Email:          user.Email,
		FullName:       user.FullName,
		Gender:         user.Gender,
		Address:        user.Address,
		PhoneNumber:    user.PhoneNumber,
		ProfilPhotoUrl: user.ProfilPhotoUrl,
		Created_at:     user.CreatedAt,
		Updated_at:     user.UpdatedAt,
	}
	for _, v := range user.Donation {
		var donation = struct {
			ID          string     `json:"id"`
			Title       string     `json:"title"`
			Description string     `json:"description"`
			PhotoUrl    string     `json:"photo_url"`
			Location    string     `json:"location"`
			Status      string     `json:"status"`
			TakerID     *string    `json:"taker_id"`
			CreatedAt   *time.Time `json:"created_at"`
			UpdatedAt   *time.Time `json:"updated_at"`
			Donator     struct {
				ID             string `json:"id"`
				UserName       string `json:"user_name"`
				FullName       string `json:"full_name"`
				PhoneNumber    string `json:"phone_number"`
				ProfilPhotoUrl string `json:"profil_photo_url"`
			} `json:"donator"`
		}{}
		donation.ID = v.ID
		donation.Title = v.Title
		donation.Description = v.Description
		donation.PhotoUrl = v.PhotoUrl
		donation.Location = v.Location
		donation.Status = v.Status
		donation.TakerID = v.TakerID
		donation.CreatedAt = v.CreatedAt
		donation.UpdatedAt = v.UpdatedAt
		donation.Donator.ID = v.User.ID
		donation.Donator.UserName = v.User.UserName
		donation.Donator.FullName = v.User.FullName
		donation.Donator.PhoneNumber = v.User.PhoneNumber
		donation.Donator.ProfilPhotoUrl = v.User.ProfilPhotoUrl
		response.Donation = append(response.Donation, donation)
	}

	return response
}
func (u UserDb) GetById(id string) models.GetUserResponse {
	User := models.User{}

	u.db.Where("id =? ", id).First(&User)
	response := models.GetUserResponse{
		ID:             User.ID,
		UserName:       User.UserName,
		Email:          User.Email,
		FullName:       User.FullName,
		Gender:         User.Gender,
		Address:        User.Address,
		PhoneNumber:    User.PhoneNumber,
		ProfilPhotoUrl: User.ProfilPhotoUrl,
		Created_at:     User.CreatedAt,
		Updated_at:     User.UpdatedAt,
	}
	return response
}

func (u UserDb) Edit(username string, newUserData models.User) (models.EditUserResponse, error) {
	User := models.User{}
	err := u.db.Model(&User).Clauses(clause.Returning{}).Where("user_name=?", username).Updates(models.User{
		Email:       newUserData.Email,
		UserName:    newUserData.UserName,
		FullName:    newUserData.FullName,
		Gender:      newUserData.Gender,
		PhoneNumber: newUserData.PhoneNumber,
		Address:     newUserData.Address,
	}).Error

	if err != nil {
		return models.EditUserResponse{}, err
	}

	response := models.EditUserResponse{
		ID:             User.ID,
		UserName:       User.UserName,
		Email:          User.Email,
		FullName:       User.FullName,
		Gender:         User.Gender,
		Address:        User.Address,
		PhoneNumber:    User.PhoneNumber,
		ProfilPhotoUrl: User.ProfilPhotoUrl,
		Updated_at:     User.UpdatedAt,
	}

	return response, nil
}

func (u UserDb) SetProfilePhoto(id, photoUrl string) (models.EditUserResponse, error) {
	User := models.User{
		ID: id,
	}
	err := u.db.Model(&User).Clauses(clause.Returning{}).Update("profil_photo_url", photoUrl).Error
	if err != nil {
		return models.EditUserResponse{}, err
	}
	response := models.EditUserResponse{
		ID:             User.ID,
		UserName:       User.UserName,
		Email:          User.Email,
		FullName:       User.FullName,
		Gender:         User.Gender,
		Address:        User.Address,
		PhoneNumber:    User.PhoneNumber,
		ProfilPhotoUrl: User.ProfilPhotoUrl,
		Updated_at:     User.UpdatedAt,
	}
	return response, nil
}

func (u UserDb) Delete(id string) error {
	User := models.User{
		ID: id,
	}
	err := u.db.Delete(&User).Error
	if err != nil {
		return err
	}
	return nil
}
