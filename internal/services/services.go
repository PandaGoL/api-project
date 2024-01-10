package services

import "github.com/PandaGoL/api-project/internal/database/postgres/models"

type UserService interface {
	AddOrUpdateUser(requestId string, user models.User) (*models.User, error)
	GetUsers(requestId string) ([]*models.User, int, error)
	GetUser(requestId string, userId string) (*models.User, error)
	DeleteUser(requestId string, userId string) error
}

type SystemService interface {
	BDCheck() error
}
