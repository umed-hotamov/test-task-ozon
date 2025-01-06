package repository

type URLRepository interface {
  Create(url string)
  Read(url string) 
}
