package model

type Books struct {
	Books []Book
}

type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"title"`
	Author string `json:"Author"`
	Owner  `json:"Owner"`
}

type Owners struct {
	Owners []Owner
}

type Owner struct {
	ID   int
	Name string
}
