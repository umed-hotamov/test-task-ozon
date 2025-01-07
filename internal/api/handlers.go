package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/umed-hotamov/url-shortener/internal/domain"
	"github.com/umed-hotamov/url-shortener/internal/service"
	"go.uber.org/zap"
)

type Handler struct {
  logger  *zap.Logger
  service *service.Service
}

func NewHandler(logger *zap.Logger, service *service.Service) *Handler {
  return &Handler{
    logger:  logger,
    service: service,
  }
}

func (h *Handler) ShortenURLHandler(c *gin.Context) {
  url := c.Param("url")
  if url == "" {
    c.JSON(http.StatusNotFound, gin.H{"Error": "url is empty"})
    h.logger.Error("error: empty url", zap.Error(errors.New("Empty url")))
    return
  }

  shortenedURL := h.service.GetShortened(domain.URL(url))
  c.JSON(http.StatusOK, gin.H{"Shortened url": fmt.Sprintf("%s", shortenedURL)})
}

func (h *Handler) OriginURLHandler(c *gin.Context) {
  url := c.Param("url")
  if url == "" {
    c.JSON(http.StatusNotFound, gin.H{"Error": "url is empty"})
    h.logger.Error("error: empty url", zap.Error(errors.New("Empty url")))
    return
  }

  originURL, err := h.service.GetOrigin(domain.URL(url))
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
    h.logger.Error("error: origin url not found", zap.Error(errors.New("Url not found")))
  }

  c.JSON(http.StatusOK, gin.H{"Origin url": fmt.Sprintf("%s", originURL)})
}
