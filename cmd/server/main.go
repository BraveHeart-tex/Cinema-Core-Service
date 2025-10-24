package main

import (
	"log"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/app"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/audit"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/config"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/db"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/handlers"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/logger"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/services"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	database, err := config.ConnectDatabase(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	db.Migrate(database)

	logger.Init()
	audit.Init(logger.Logger)

	// Setup DI
	userRepo := repositories.NewUserRepository(database)
	sessionRepo := repositories.NewSessionRepository(database)
	sessionService := services.NewSessionService(sessionRepo)
	userService := services.NewUserService(userRepo, sessionService)
	userHandler := handlers.NewUserHandler(userService)

	appCtx := &app.App{
		UserHandler: userHandler,
	}

	router := app.SetupRouter(appCtx)
	router.Run()
}
