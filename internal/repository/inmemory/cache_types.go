package inmemory

type ID  int
type URL string

func (id ID) Inc() {
  id += 1
}

func (id ID) Reset() {
  id = 0
}

func StringToURL(url string) URL {
  return URL(url) 
}

func INToID(id int) ID {
  return ID(id)
}
