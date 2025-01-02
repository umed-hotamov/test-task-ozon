package main

import (
	"database/sql"
	"log"

	"github.com/umed-hotamov/url-shortener/internal/config"
	"github.com/umed-hotamov/url-shortener/pkg/database"
	"github.com/umed-hotamov/url-shortener/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg := config.GetConfig()
	logger, err := logger.NewLogger("info")
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}

	db, err := database.InitDB(cfg, logger)
	if err != nil {
		logger.Fatal("Failed to create database")
	}
	defer db.Close()

	run(logger, db)
}

func run(logger *zap.Logger, db *sql.DB) {
	router := gin.Default()
}
