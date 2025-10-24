package config

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DatabaseURL string
	ServerPort  string
	AppEnv      string
}

func LoadConfig() (*Config, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		fmt.Println("APP_ENV is not set, defaulting to development")
		appEnv = "development"
	}

	if appEnv != "development" && appEnv != "production" {
		return nil, fmt.Errorf("invalid APP_ENV value: %s", appEnv)
	}

	return &Config{
		DatabaseURL: dbURL,
		ServerPort:  port,
		AppEnv:      appEnv,
	}, nil
}

func ConnectDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
