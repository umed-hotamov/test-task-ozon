package domain 

type ID  int
type URL string

func (id ID) Inc() {
  id += 1
}

func (id ID) Reset() {
  id = 0
}

func (url URL) ToString() string {
  return string(url)
}

func (id ID) ToInt() int {
  return int(id)
}
