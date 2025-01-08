package main

import (
	"fmt"
	"log"

	handler "github.com/umed-hotamov/url-shortener/internal/api"
	"github.com/umed-hotamov/url-shortener/internal/config"
	"github.com/umed-hotamov/url-shortener/internal/repository"
	database "github.com/umed-hotamov/url-shortener/internal/repository/inmemory"
	"github.com/umed-hotamov/url-shortener/internal/service"
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
  handler := handler.NewHandler(logger, service)

  r.POST("/shorten", handler.ShortenURLHandler)
  r.GET("/origin", handler.OriginURLHandler)

  r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
