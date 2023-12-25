package user

import (
	"time"

	"github.com/PandaGoL/api-project/internal/database"
	"github.com/PandaGoL/api-project/internal/database/postgres/models"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	s database.Storage
}

type User struct {
	UserID    string
	FirstName string
	LastName  string
	Age       int
	Email     string
	Phone     string
}

func NewUserService(db database.Storage) *UserService {
	us := &UserService{}
	us.s = db
	return us
}

func (u *User) ToMap() map[string]interface{} {
	fields := make(map[string]interface{})
	fields["id"] = u.UserID
	fields["first_name"] = u.FirstName
	fields["last_name"] = u.LastName
	fields["email"] = u.Email
	fields["phone"] = u.Phone

	return fields
}

func (us *UserService) AddOrUpdateUser(user models.User) (*models.User, error) {
	bt := time.Now()
	logrus.WithField("time_start", bt).WithFields(user.ToMap()).Info("Start AddorUpdateUser request: ")
	resp := new(models.User)
	resp, err := us.s.AddOrUpdateUser(user)
	if err != nil {
		logrus.Errorf("Error in AddOrUpdateUser: %s", err)
		return resp, err
	}
	logrus.WithField("time_finish", time.Now()).WithField("durations", time.Since(bt)).Info("Finish AddorUpdateUser request")
	return resp, nil
}

func (us *UserService) GetUsers() ([]*models.User, int, error) {
	bt := time.Now()
	logrus.WithField("time_start", bt).Info("Start GetUsers request")
	resp, count, err := us.s.GetUsers()
	if err != nil {
		logrus.Errorf("Error in GetUsers: %s", err)
		return resp, 0, err
	}
	logrus.WithField("time_finish", time.Now()).WithField("durations", time.Since(bt)).Info("Finish GetUsers request")
	return resp, count, nil
}

func (us *UserService) GetUser(userId string) (*models.User, error) {

	bt := time.Now()
	logrus.WithField("time_start", bt).WithField("userId", userId).Info("Start GetUser request")
	resp, err := us.s.GetUser(userId)
	if err != nil {
		logrus.Errorf("Error in GetUser: %s", err)
		return resp, err
	}
	logrus.WithField("time_finish", time.Now()).WithField("durations", time.Since(bt)).Info("Finish GetUser request")
	return resp, nil
}

func (us *UserService) DeleteUser(userId string) error {
	bt := time.Now()
	logrus.WithField("time_start", bt).Info("Start DeleteUser request")
	err := us.s.DeleteUser(userId)
	if err != nil {
		logrus.Errorf("Error in DeleteUser: %s", err)
		return err
	}
	logrus.WithField("time_finish", time.Now()).WithField("durations", time.Since(bt)).Info("Finish DeleteUser request")
	return nil
}
