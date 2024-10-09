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
	ctx := r.Context()
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	err := h.serv.CreateUser(ctx, &user)
	if err != nil {
		commonHandler.Error(w, err)
		return
	}

	commonHandler.Respond(w, http.StatusCreated, user)
}

func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := h.serv.GetUserByID(ctx, id)
	if err != nil {
		commonHandler.Error(w, err)
		return
	}

	commonHandler.Respond(w, http.StatusOK, user)
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	err := h.serv.UpdateUser(ctx, &user)
	if err != nil {
		commonHandler.Error(w, err)
		return
	}

	commonHandler.Respond(w, http.StatusOK, nil)
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.serv.DeleteUser(ctx, id)
	if err != nil {
		commonHandler.Error(w, err)
		return
	}

	commonHandler.Respond(w, http.StatusOK, nil)
}

func (h *userHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := h.serv.GetAllUsers(ctx)
	if err != nil {
		commonHandler.Error(w, err)
		return
	}

	commonHandler.Respond(w, http.StatusOK, users)
}
