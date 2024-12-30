package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/umed-hotamov/url-shortener/pkg/logger"
	"github.com/umed-hotamov/url-shortener/config"
)

func Run(cfg *config.Config, db *sql.DB) {
  router := gin.Default()
  
  logger, _ := logger.NewLogger("info")
  
}
