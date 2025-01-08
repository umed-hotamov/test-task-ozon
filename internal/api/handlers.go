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

type Request struct {
  URL string `json:"url"`
}

func (h *Handler) ShortenURLHandler(c *gin.Context) {
  var req Request
  
  if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
    h.logger.Error("error: bad request", zap.Error(err))
  }
  
  if req.URL == "" {
    c.JSON(http.StatusNotFound, gin.H{"Error": "url is empty"})
    h.logger.Error("error: empty url", zap.Error(errors.New("Empty url")))
    return
  }

  shortenedURL := h.service.GetShortened(domain.URL(req.URL))
  c.JSON(http.StatusOK, gin.H{"Shortened url": fmt.Sprintf("%s", shortenedURL)})
}

func (h *Handler) OriginURLHandler(c *gin.Context) {
  var req Request
  
  if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
    h.logger.Error("error: bad request", zap.Error(err))
  }

  if req.URL == "" {
    c.JSON(http.StatusNotFound, gin.H{"Error": "url is empty"})
    h.logger.Error("error: empty url", zap.Error(errors.New("Empty url")))
    return
  }

  originURL, err := h.service.GetOrigin(domain.URL(req.URL))
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
    h.logger.Error("error: origin url not found", zap.Error(errors.New("Url not found")))
  }

  c.JSON(http.StatusOK, gin.H{"Origin url": fmt.Sprintf("%s", originURL)})
}
