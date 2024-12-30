package service

import (

)

type Url struct {
  ID        int    `json:"id,omitempty"`
  OriginUrl string `json:"origin_url"`
  ShortUrl  string `json:"short_url,omitempty"`
}

func (s *Service) AddUrl(url *Url) {
  const query = `INSERT INTO urls (origin_url)
                 VALUES ($1)`


} 
