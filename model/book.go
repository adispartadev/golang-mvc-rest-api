package model

type BookList struct {
	Books []Book
}

type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"title"`
	Author string `json:"Author"`
	Owner  `json:"Owner"`
}

type OwnerList struct {
	Owners []Owner
}

type Owner struct {
	ID   int
	Name string
}
