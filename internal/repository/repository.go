package repository

import (
  "github.com/umed-hotamov/url-shortener/internal/domain"
)

type URLRepository interface {
  Set(url domain.URL)
  Get(id domain.ID) (domain.URL, error)
  LastID() domain.ID
}
