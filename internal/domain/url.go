package domain

type ID  int
type URL string 

func (url URL) ToString() string {
  return string(url)
}

func (id ID) ToInt() int {
  return int(id)
}
