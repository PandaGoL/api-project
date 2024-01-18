package user

import (
	"net/http"
	"time"

	"github.com/PandaGoL/api-project/internal/database"
	"github.com/PandaGoL/api-project/internal/database/postgres/models"
	"github.com/PandaGoL/api-project/internal/metrics"
	"github.com/PandaGoL/api-project/pkg/recovery"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	s        database.Storage
	m        metrics.Metrics
	recovery recovery.Recovery
}

type User struct {
	UserID    string
	FirstName string
	LastName  string
	Age       int
	Email     string
	Phone     string
}

func NewUserService(db database.Storage, m metrics.Metrics, r recovery.Recovery) *UserService {
	us := &UserService{
		s:        db,
		m:        m,
		recovery: r,
	}
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

func (us *UserService) AddOrUpdateUser(requestId string, user models.User) (*models.User, error) {
	bt := time.Now()
	us.m.AddHTTPIncomingRequests("/v1/api/adduser")

	logrus.WithField("request_id", requestId).WithField("time_start", bt).WithFields(user.ToMap()).Info("Start AddorUpdateUser request: ")

	defer us.recovery.Do()

	resp := new(models.User)
	resp, err := us.s.AddOrUpdateUser(user)
	if err != nil {
		us.m.AddHTTPIncomingResponses("/v1/api/adduser", http.StatusInternalServerError, time.Since(bt))
		logrus.Errorf("request_id: %s Error in AddOrUpdateUser: %s", requestId, err)
		return resp, err
	}
	us.m.AddHTTPIncomingResponses("/v1/api/adduser", http.StatusOK, time.Since(bt))
	logrus.WithField("request_id", requestId).WithField("time_finish", time.Now()).WithField("durations", time.Since(bt)).Info("Finish AddorUpdateUser request")
	return resp, nil
}

func (us *UserService) GetUsers(requestId string) ([]*models.User, int, error) {
	bt := time.Now()
	us.m.AddHTTPIncomingRequests("/v1/api/users")
	logrus.WithField("request_id", requestId).WithField("time_start", bt).Info("Start GetUsers request")

	defer us.recovery.Do()

	resp, count, err := us.s.GetUsers()
	if err != nil {
		us.m.AddHTTPIncomingResponses("/v1/api/users", http.StatusInternalServerError, time.Since(bt))
		logrus.Errorf("request_id: %s Error in GetUsers: %s", requestId, err)
		return resp, 0, err
	}
	us.m.AddHTTPIncomingResponses("/v1/api/users", http.StatusOK, time.Since(bt))
	logrus.WithField("request_id", requestId).WithField("time_finish", time.Now()).WithField("durations", time.Since(bt)).Info("Finish GetUsers request")
	return resp, count, nil
}

func (us *UserService) GetUser(requestId string, userId string) (*models.User, error) {

	bt := time.Now()
	us.m.AddHTTPIncomingRequests("/v1/api/user/{user_id}")
	logrus.WithField("request_id", requestId).WithField("time_start", bt).WithField("userId", userId).Info("Start GetUser request")

	defer us.recovery.Do()

	resp, err := us.s.GetUser(userId)
	if err != nil {
		us.m.AddHTTPIncomingResponses("/v1/api/user/{user_id}", http.StatusInternalServerError, time.Since(bt))
		logrus.Errorf("request_id: %s Error in GetUser: %s", requestId, err)
		return resp, err
	}
	us.m.AddHTTPIncomingResponses("/v1/api/user/{user_id}", http.StatusOK, time.Since(bt))
	logrus.WithField("request_id", requestId).WithField("time_finish", time.Now()).WithField("durations", time.Since(bt)).Info("Finish GetUser request")
	return resp, nil
}

func (us *UserService) DeleteUser(requestId string, userId string) error {
	bt := time.Now()
	us.m.AddHTTPIncomingRequests("/v1/api/user")
	logrus.WithField("request_id", requestId).WithField("time_start", bt).Info("Start DeleteUser request")

	defer us.recovery.Do()

	err := us.s.DeleteUser(userId)
	if err != nil {
		us.m.AddHTTPIncomingResponses("/v1/api/user/{user_id}", http.StatusInternalServerError, time.Since(bt))
		logrus.Errorf("request_id: %s Error in DeleteUser: %s", requestId, err)
		return err
	}
	us.m.AddHTTPIncomingResponses("/v1/api/user/{user_id}", http.StatusOK, time.Since(bt))
	logrus.WithField("request_id", requestId).WithField("time_finish", time.Now()).WithField("durations", time.Since(bt)).Info("Finish DeleteUser request")
	return nil
}
