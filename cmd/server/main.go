package main

import (
	"log"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/app"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/audit"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/config"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/db"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/handlers"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/handlers/admin"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/logger"
	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/repositories"
	adminServices "github.com/BraveHeart-tex/Cinema-Core-Service/internal/services/admin"
	sessionServices "github.com/BraveHeart-tex/Cinema-Core-Service/internal/services/session"
	userServices "github.com/BraveHeart-tex/Cinema-Core-Service/internal/services/user"
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

	// Initialize transaction manager
	txManager := db.NewGormTxManager(database)

	// ================= Repositories =================
	userRepo := repositories.NewUserRepository(database)
	sessionRepo := repositories.NewSessionRepository(database)
	genreRepo := repositories.NewGenreRepository(database)
	movieRepo := repositories.NewMovieRepository(database)
	theaterRepo := repositories.NewTheaterRepository(database)
	showtimeRepo := repositories.NewShowtimeRepository(database)

	// ================= Services =================
	sessionService := sessionServices.NewSessionService(sessionRepo)
	userService := userServices.NewUserService(userRepo, sessionService, txManager)

	// Admin services - each domain gets its own service
	// Dependency flow: Repositories -> Domain Services -> Aggregator
	adminServices := adminServices.NewServices(userRepo, genreRepo, movieRepo, theaterRepo, showtimeRepo)

	// ================= Handlers =================
	userHandler := handlers.NewUserHandler(userService)

	// Admin handler aggregates all domain services
	// This single handler has access to all domain operations via adminServices
	adminHandler := admin.NewAdminHandler(adminServices)

	appCtx := &app.App{
		UserHandler:    userHandler,
		AdminHandler:   adminHandler,
		SessionService: sessionService,
		UserService:    userService,
	}

	router := app.SetupRouter(appCtx)
	router.Run()
}
