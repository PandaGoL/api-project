package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PandaGoL/api-project/internal/services/student"
	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Service *student.Service
}

func NewHandler(service *student.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) InitRoutes() {
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/students", h.GetAllStudents).Methods("GET")
	h.Router.HandleFunc("/api/students", h.PostStudent).Methods("POST")
	h.Router.HandleFunc("/api/students/{school}", h.GetStudentsBySchool).Methods("GET")
	h.Router.HandleFunc("/api/students/{id}", h.GetStudentByID).Methods("GET")
	h.Router.HandleFunc("api/students/{id}", h.UpdateStudent).Methods("PUT")
	h.Router.HandleFunc("/api/students/{id}", h.DeleteStudent).Methods("DELETE")
	h.Router.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Status Up!")
	})
}

func (h *Handler) GetAllStudents(w http.ResponseWriter, r *http.Request)      {}
func (h *Handler) PostStudent(w http.ResponseWriter, r *http.Request)         {}
func (h *Handler) GetStudentsBySchool(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	
	vars := mux.Vars(r)
	id := vars["id"]

	studentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		respondWithError(w, "Error Parsing ID to UINT", err)
	}

	student, err := h.Service.GetStudentByID(uint(studentID))
	if err != nil {
		respondWithError(w, "Error Retriening Student by ID", err)
	}

	if err := json.NewEncoder(w).Encode(student); err != nil {
		panic(err)
	}
}
func (h *Handler) UpdateStudent(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) DeleteStudent(w http.ResponseWriter, r *http.Request) {}
