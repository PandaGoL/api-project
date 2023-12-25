package v1

import (
	"encoding/json"
	"net/http"

	"github.com/PandaGoL/api-project/internal/api/http/types"
	"github.com/PandaGoL/api-project/internal/database/postgres/models"
	"github.com/sirupsen/logrus"
)

func (rout *Router) AddOrUpdateUser(w http.ResponseWriter, r *http.Request) {

	rec := models.User{}
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	resp, err := rout.Service.AddOrUpdateUser(rec)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		logrus.Errorf("Unable to encode json: %v\n", err)
		w.WriteHeader(500)
		return
	}
}

func (rout *Router) GetUsers(w http.ResponseWriter, r *http.Request) {

	resp, count, err := rout.Service.GetUsers()
	if err != nil {
		w.WriteHeader(400)
		return
	}

	res := &types.GetUserResponse{
		User:  resp,
		Count: count,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		logrus.Errorf("Unable to encode json: %v\n", err)
		w.WriteHeader(500)
		return
	}
}

func (rout *Router) GetUser(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("user_id")
	resp, err := rout.Service.GetUser(id)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		logrus.Errorf("Unable to encode json: %v\n", err)
		w.WriteHeader(500)
		return
	}
}

func (rout *Router) DeleteUser(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("user_id")
	err := rout.Service.DeleteUser(id)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(err)
	if err != nil {
		logrus.Errorf("Unable to encode json: %v\n", err)
		w.WriteHeader(500)
		return
	}
}
