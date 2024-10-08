package router

import (
	"github.com/Beretta350/golang-rest-template/internal/app/common/middleware"
	"github.com/Beretta350/golang-rest-template/internal/app/user/handler"
	"github.com/gorilla/mux"
)

const (
	UsersEnpointPath       = "/users"
	UsersEnpointWithIDPath = "/users/{id}"
)

func Router(userHandler handler.UserHandler) *mux.Router {

	mux := mux.NewRouter()

	mux.Use(middleware.LoggingMiddleware)

	mux.HandleFunc(UsersEnpointPath, userHandler.CreateUser).Methods("POST")
	mux.HandleFunc(UsersEnpointPath, userHandler.GetAllUsers).Methods("GET")
	mux.HandleFunc(UsersEnpointWithIDPath, userHandler.GetUser).Methods("GET")
	mux.HandleFunc(UsersEnpointWithIDPath, userHandler.UpdateUser).Methods("PUT")
	mux.HandleFunc(UsersEnpointWithIDPath, userHandler.DeleteUser).Methods("DELETE")

	return mux
}
