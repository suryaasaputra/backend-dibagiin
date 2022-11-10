package repository

import (
	"dibagi/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUserRepository interface {
	RegisterUser(models.User) (models.CreateUserResponse, error)
	EditUser(string, models.User) (models.EditUserResponse, error)
	GetUserByEmail(string) models.User
	GetUserByUserName(string) models.GetUserResponse
	SetProfilePhoto(string, string) error
	DeleteUser(string) error
}

type UserDb struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserDb {
	return &UserDb{
		db: db,
	}
}

func (u UserDb) RegisterUser(user models.User) (models.CreateUserResponse, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return models.CreateUserResponse{}, err
	}

	return models.CreateUserResponse{
		ID:             user.ID,
		UserName:       user.UserName,
		Email:          user.Email,
		FullName:       user.FullName,
		Age:            user.Age,
		Gender:         user.Gender,
		PhoneNumber:    user.PhoneNumber,
		Address:        user.Address,
		ProfilPhotoUrl: user.ProfilPhotoUrl,
		CreatedAt:      user.CreatedAt,
	}, nil
}

func (u UserDb) GetUserByEmail(email string) models.User {
	User := models.User{}

	u.db.Where("email =? ", email).First(&User)

	return User
}

func (u UserDb) GetUserByUserName(userName string) models.GetUserResponse {
	User := models.User{}

	u.db.Preload("Donation").Where("user_name =? ", userName).First(&User)
	response := models.GetUserResponse{
		ID:             User.ID,
		UserName:       User.UserName,
		Email:          User.Email,
		FullName:       User.FullName,
		Age:            User.Age,
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

func (u UserDb) EditUser(username string, newUserData models.User) (models.EditUserResponse, error) {
	User := models.User{}
	err := u.db.Model(&User).Clauses(clause.Returning{}).Where("user_name=?", username).Updates(models.User{
		Email:       newUserData.Email,
		UserName:    newUserData.UserName,
		FullName:    newUserData.FullName,
		Age:         newUserData.Age,
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
		Age:            User.Age,
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

func (u UserDb) DeleteUser(id string) error {
	User := models.User{
		ID: id,
	}
	err := u.db.Delete(&User).Error
	if err != nil {
		return err
	}
	return nil
}
