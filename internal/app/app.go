package app

import (
	"log"
	"net/http"

	"github.com/Beretta350/golang-rest-template/config"
	"github.com/Beretta350/golang-rest-template/internal/app/common/router"
	"github.com/Beretta350/golang-rest-template/internal/app/user/handler"
	"github.com/Beretta350/golang-rest-template/internal/app/user/repository"
	"github.com/Beretta350/golang-rest-template/internal/app/user/service"
	"github.com/Beretta350/golang-rest-template/internal/pkg/database"
)

func Run(env string) {
	log.Println("Environment:", env)
	if env == "" {
		env = "local"
	}

	cfg, err := config.LoadConfig(env)
	if err != nil {
		log.Fatalln("Error loading config:", err.Error())
	}

	_, mongodb, err := database.InitDatabase(&cfg.Database)
	if err != nil {
		log.Fatalln("Error establishing connection to database:", err.Error())
	}

	userRepo := repository.NewUserRepository(mongodb)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	//routes
	mux := router.Router(userHandler)

	//run
	log.Printf("Server running on port %v\n", cfg.Server.Port)
	err = http.ListenAndServe(":"+cfg.Server.Port, mux)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
