package inmemory

import (
  "sync"
  "time"
)

var id ID

type URLCache struct {
  items map[ID]URL
  mu    sync.Mutex
}

func NewURLCache() *URLCache {
  c := &URLCache{
    items: make(map[ID]URL),
  }

  go func() {
    for range time.Tick(5 * time.Hour) {
      c.mu.Lock()
      id.Reset()
      c.mu.Unlock()
    }
  }()

  return c
}

func (uc *URLCache) Create(url string) {
  uc.mu.Lock()
  defer uc.mu.Unlock()
  
  uc.items[id] = StringToURL(url)
  id.Inc()
}

func (uc *URLCache) Read(id int) (URL, bool) {
  uc.mu.Lock()
  defer uc.mu.Unlock()

  url, found := uc.items[INToID(id)]

  return url, found
}
