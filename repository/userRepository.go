package repository

import (
	"dibagi/models"
	"errors"

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
	SetProfilePhoto(string, string) error
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
	User := models.User{}

	u.db.Preload("Donation").Where("user_name =? ", userName).First(&User)
	response := models.GetUserResponse{
		ID:             User.ID,
		UserName:       User.UserName,
		Email:          User.Email,
		FullName:       User.FullName,
		Gender:         User.Gender,
		Address:        User.Address,
		PhoneNumber:    User.PhoneNumber,
		ProfilPhotoUrl: User.ProfilPhotoUrl,
		Donation:       User.Donation,
		Created_at:     User.CreatedAt,
		Updated_at:     User.UpdatedAt,
	}
	return response
}
func (u UserDb) GetById(id string) models.GetUserResponse {
	User := models.User{}

	u.db.Preload("Donation").Where("id =? ", id).First(&User)
	response := models.GetUserResponse{
		ID:             User.ID,
		UserName:       User.UserName,
		Email:          User.Email,
		FullName:       User.FullName,
		Gender:         User.Gender,
		Address:        User.Address,
		PhoneNumber:    User.PhoneNumber,
		ProfilPhotoUrl: User.ProfilPhotoUrl,
		Donation:       User.Donation,
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

func (u UserDb) SetProfilePhoto(id, photoUrl string) error {
	User := models.User{
		ID: id,
	}
	err := u.db.Model(&User).Update("profil_photo_url", photoUrl).Error
	if err != nil {
		return err
	}
	return nil
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
