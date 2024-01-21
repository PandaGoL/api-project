package v1

import (
	"net/http"

	"github.com/PandaGoL/api-project/internal/services"
	"github.com/gorilla/mux"
)

// Router - cтруктура данных для HTTP API v1
type Router struct {
	mainRouter *mux.Router
	router     *mux.Router
	Service    services.UserService
}

// InitAPI - функция инициализирует HTTP API версии 1
func InitAPI(mainRouter *mux.Router, s services.UserService, middleware func(next http.Handler) http.Handler) {
	sr := &Router{
		mainRouter: mainRouter,
		router:     mainRouter.PathPrefix("/v1").Subrouter(),
		Service:    s,
	}
	sr.router.Use(middleware)

	sr.router.HandleFunc("/api/adduser", sr.AddOrUpdateUser).Methods("POST")
	sr.router.HandleFunc("/api/users", sr.GetUsers).Methods("Get")
	sr.router.HandleFunc("/api/user/{user_id}", sr.GetUser).Methods("Get")
	sr.router.HandleFunc("/api/user", sr.DeleteUser).Methods("Delete")
}
