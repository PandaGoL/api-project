package database

import "github.com/PandaGoL/api-project/internal/database/postgres/models"

type Storage interface {
	AddOrUpdateUser(user models.User) (scanUser *models.User, err error)
	GetUsers() (users []*models.User, count int, err error)
	GetUser(userID string) (user *models.User, err error)
	DeleteUser(userID string) error
	Close()
}
