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

	// ================= Repositories =================
	userRepo := repositories.NewUserRepository(database)
	sessionRepo := repositories.NewSessionRepository(database)

	// ================= Services =================
	sessionService := services.NewSessionService(sessionRepo)
	userService := services.NewUserService(userRepo, sessionService)
	adminService := services.NewAdminService(userRepo)

	// ================= Handlers =================
	userHandler := handlers.NewUserHandler(userService)
	adminHandler := handlers.NewAdminHandler(adminService)

	appCtx := &app.App{
		AdminHandler:   adminHandler,
		UserHandler:    userHandler,
		SessionService: sessionService,
		UserService:    userService,
	}

	router := app.SetupRouter(appCtx)
	router.Run()
}
