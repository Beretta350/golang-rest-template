package app

import (
	"net/http"

	"github.com/Beretta350/golang-rest-template/config"
	"github.com/Beretta350/golang-rest-template/internal/app/common/router"
	"github.com/Beretta350/golang-rest-template/internal/app/user/handler"
	"github.com/Beretta350/golang-rest-template/internal/app/user/repository"
	"github.com/Beretta350/golang-rest-template/internal/app/user/service"
	"github.com/Beretta350/golang-rest-template/internal/pkg/database"
)

func Run(env string) {
	if env == "" {
		env = "local"
	}

	cfg, err := config.LoadConfig(env)
	if err != nil {
		panic("Error loading config!")
	}

	_, mongodb, err := database.InitDatabase(&cfg.Database)
	if err != nil {
		panic("Error establishing connection to database")
	}

	userRepo := repository.NewUserRepository(mongodb)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	mux := router.Router(userHandler)
	http.ListenAndServe(cfg.Server.Port, mux)
}
