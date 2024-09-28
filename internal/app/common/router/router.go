package router

import (
	"github.com/Beretta350/golang-rest-template/internal/app/user/handler"
	"github.com/gorilla/mux"
)

func Router(userHandler handler.UserHandler) *mux.Router {

	mux := mux.NewRouter()
	mux.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	mux.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	mux.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	mux.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	mux.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	return mux
}
