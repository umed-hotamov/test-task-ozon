package inmemory

import (
	"errors"
	"sync"
	"time"

	"github.com/umed-hotamov/url-shortener/internal/domain"
)

var id domain.ID = 1

type URLCache struct {
  items map[domain.ID]domain.URL
  mu    sync.Mutex
}

func NewURLCache() *URLCache {
  c := &URLCache{
    items: make(map[domain.ID]domain.URL),
  }

  go func() {
    for range time.Tick(5 * time.Hour) {
      c.mu.Lock()
      clear(c.items)
      id = 0
      c.mu.Unlock()
    }
  }()

  return c
}

func (uc *URLCache) Set(url domain.URL) {
  uc.mu.Lock()
  defer uc.mu.Unlock()
  
  uc.items[id] = url
  id += 1
}

func (uc *URLCache) Get(id domain.ID) (domain.URL, error) {
  uc.mu.Lock()
  defer uc.mu.Unlock()

  url, found := uc.items[id]

  var err error
  if found == false {
    err = errors.New("url does not exist")
  } 

  return url, err
}

func (uc *URLCache) LastID() domain.ID {
  return id - 1
}
