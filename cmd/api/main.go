package main

import (
	"log"

	"github.com/umed-hotamov/url-shortener/internal/config"
	"github.com/umed-hotamov/url-shortener/internal/service"
  database "github.com/umed-hotamov/url-shortener/internal/repository/inmemory"
	"github.com/umed-hotamov/url-shortener/internal/repository"
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
	db := database.NewURLCache()

	run(cfg, logger, db)
}

func run(cfg *config.Config, logger *zap.Logger, db repository.URLRepository) {
	r := gin.Default()

  service := service.NewService(db) 
  handler := NewHandler(logger, service)

  r.GET("shorten/:url", handler.ShortenURLHandler)
  r.GET("origin/:url", handler.OriginURLHandler)

  r.Run(cfg.Server.Port)
}
