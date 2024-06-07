package storage

import "github.com/TimurZheksimbaev/Golang-auth/models"

type UserStorageI interface {
	Create(*models.User) error
	Update(models.User) error
	Delete(models.User) error
	FindByID(uint) (models.User, error)
	FindByEmail(string) (*models.User, error)
	FindByUsername(string) (models.User, error)
	GetAll() ([]models.User, error)
}