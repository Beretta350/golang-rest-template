package handler

import (
	"encoding/json"
	"net/http"

	commonHandler "github.com/Beretta350/golang-rest-template/internal/app/common/handler"
	"github.com/Beretta350/golang-rest-template/internal/app/user/model"
	"github.com/Beretta350/golang-rest-template/internal/app/user/service"
	"github.com/gorilla/mux"
)

// UserHandler defines the HTTP handlers for user-related operations.
type UserHandler interface {
	// CreateUser handles the HTTP request for creating a new user.
	// It reads user data from the request body and responds with the created user.
	CreateUser(w http.ResponseWriter, r *http.Request)

	// GetUser handles the HTTP request to retrieve a user by their unique identifier (ID).
	// It expects the user ID to be provided in the URL and responds with the user data.
	GetUser(w http.ResponseWriter, r *http.Request)

	// UpdateUser handles the HTTP request for updating an existing user's information.
	// It reads the updated user data from the request body and responds with the updated user details.
	UpdateUser(w http.ResponseWriter, r *http.Request)

	// DeleteUser handles the HTTP request to remove a user by their unique identifier (ID).
	// It expects the user ID to be provided in the URL and responds with a success or error message.
	DeleteUser(w http.ResponseWriter, r *http.Request)

	// GetAllUsers handles the HTTP request to retrieve a list of all users.
	// It responds with an array of user data.
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

	vars := mux.Vars(r)
	id := vars["id"]

	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	user.Id = id

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
