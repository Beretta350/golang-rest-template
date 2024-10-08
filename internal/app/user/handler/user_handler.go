package handler

import (
	"encoding/json"
	"net/http"

	commonHandler "github.com/Beretta350/golang-rest-template/internal/app/common/handler"
	"github.com/Beretta350/golang-rest-template/internal/app/user/model"
	"github.com/Beretta350/golang-rest-template/internal/app/user/service"
	"github.com/gorilla/mux"
)

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	serv service.UserService
}

func NewUserHandler(s service.UserService) UserHandler {
	return &userHandler{serv: s}
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	err := h.serv.Create(r.Context(), &user)
	if err != nil {
		commonHandler.HttpError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := h.serv.GetByID(r.Context(), id)
	if err != nil {
		commonHandler.HttpError(w, err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	err := h.serv.Update(r.Context(), &user)
	if err != nil {
		commonHandler.HttpError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.serv.Delete(r.Context(), id)
	if err != nil {
		commonHandler.HttpError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *userHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.serv.GetAll(r.Context())
	if err != nil {
		commonHandler.HttpError(w, err)
		return
	}

	json.NewEncoder(w).Encode(users)
}
