package main

import (
	"log"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/app"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/config"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/handlers"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/services"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := config.ConnectDatabase(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	// Setup DI
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	appCtx := &app.App{
		UserHandler: userHandler,
	}

	router := app.SetupRouter(appCtx)
	router.Run()
}
