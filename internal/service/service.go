package service

import (
  repo "github.com/umed-hotamov/url-shortener/internal/repository"
)

type Service struct {
	DB repo.URLRepository
}

func NewService(db repo.URLRepository) *Service {
	return &Service{
		DB: db,
	}
}
