package model

type OwnerList struct {
	Owners []Owner `json:"owners"`
}

type Owner struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
