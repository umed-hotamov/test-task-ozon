package repository

import (
  "github.com/umed-hotamov/url-shortener/internal/domain"
)

type URLRepository interface {
  Set(url domain.URL) domain.ID
  Get(id domain.ID) (domain.URL, error)
  LastID() domain.ID
}
