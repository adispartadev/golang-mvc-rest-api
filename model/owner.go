package model

import (
	"database/sql"
	"errors"
	e "golang-mvc-rest-api/entity"

	"golang-mvc-rest-api/db"
)

type Owner struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetAllOwners(pagination e.ResponsePagination) ([]Owner, error) {
	var owner Owner
	var ownerList []Owner

	rows, err := db.DB.Query(`SELECT * FROM owners order by id desc offset $1 limit $2`, pagination.Offset, pagination.ItemInPage)
	if err != nil {
		return ownerList, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&owner.ID, &owner.Name)
		if err != nil {
			return ownerList, err
		}
		ownerList = append(ownerList, owner)
	}

	err = rows.Err()
	if err != nil {
		return ownerList, err
	}
	return ownerList, nil
}

func CountOwners() (int, error) {
	var countOwners int

	err := db.DB.QueryRow("SELECT count(*) FROM owners").Scan(&countOwners)

	if err == sql.ErrNoRows {
		return 0, errors.New("owners is empty")
	}

	if err != nil {
		return 0, err
	}

	return countOwners, nil
}
