package storage

import (
	"github.com/TimurZheksimbaev/Golang-auth/models"
	"github.com/TimurZheksimbaev/Golang-auth/utils"
	"gorm.io/gorm"
)


type UserStorage struct {
	DB *gorm.DB
}

func NewStorage(gormDB *gorm.DB) UserStorageI {
	return &UserStorage{
		DB: gormDB,
	}
}

func (s *UserStorage) Create(user *models.User) error {
	result := s.DB.Save(&user)

	if result.Error != nil {
		return utils.DatabaseError("Could not create user", result.Error)
	}
	return nil
}

func (s *UserStorage) Update(user models.User) error {
	updateUser := models.User{
		Username: user.Username,
		Email: user.Email,
		Password: user.Password,
	}

	result := s.DB.Model(&user).Updates(updateUser)
	if result.Error != nil {
		return utils.DatabaseError("Could not update user", result.Error)
	}
	return nil
}

func (s *UserStorage) Delete(user models.User) error {
	result := s.DB.Delete(&user).Where("id = ?", user.ID)
	if result.Error != nil {
		return utils.DatabaseError("Could not delete user", result.Error)
	}
	return nil
}

func (s *UserStorage) GetAll() ([]models.User, error) {
	var users []models.User
	result := s.DB.Find(&users)
	if result.Error != nil {
		return nil, utils.DatabaseError("Could not get all users", result.Error)
	}
	return users, nil
}

func (s *UserStorage) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := s.DB.First(&user, "email=?", email)
	if result.Error != nil {
		return &models.User{}, utils.DatabaseError("Could find user with such email", result.Error)
	}
	return &user, nil
}

func (s *UserStorage) FindByUsername(username string) (models.User, error) {
	var user models.User
	result := s.DB.First(&user).Where("username = ?", username)
	if result.Error != nil {
		return models.User{}, utils.DatabaseError("Could not find user with such username", result.Error)
	}
	return user, nil
}

func (s *UserStorage) FindByID(id uint) (models.User, error) {
	var user models.User
	result := s.DB.First(&user).Where("id = ?", id)
	if result.Error != nil {
		return models.User{}, utils.DatabaseError("Could not find user with such id", result.Error)
	}
	return user, nil
}

