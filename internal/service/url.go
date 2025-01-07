package service

import (
	"github.com/umed-hotamov/url-shortener/internal/domain"
	us "github.com/umed-hotamov/url-shortener/internal/urlshortener"
)

func (s *Service) GetShortened(url domain.URL) domain.URL {
  s.DB.Set(url)

  shortenedURL := us.Encode(s.DB.LastID().ToInt())
  return domain.URL(shortenedURL)
}

func (s *Service) GetOrigin(shortenedURL domain.URL) (domain.URL, error) {
  originID := us.Decode(shortenedURL.ToString())

  originURL, err := s.DB.Get(domain.ID(originID))
  if err != nil {
    return "", err
  }

  return originURL, nil
}
