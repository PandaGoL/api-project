package services

import "github.com/PandaGoL/api-project/internal/database/postgres/models"

type UserService interface {
	AddOrUpdateUser(user models.User) (*models.User, error)
	GetUsers() ([]*models.User, int, error)
	GetUser(userId string) (*models.User, error)
	DeleteUser(userId string) error
}
