package model

import (
	"database/sql"
	"errors"
	"fmt"
	e "golang-mvc-rest-api/entity"

	"golang-mvc-rest-api/db"
)

func GetAllOwners(pagination e.ResponsePagination) ([]e.Owner, error) {
	var owner e.Owner
	var ownerList []e.Owner

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

func InsertOwner(owner *e.Owner) error {
	const query = `INSERT INTO owners (name) VALUES ($1)`
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	res, err := tx.Exec(query, owner.Name)
	if err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println("=======print res======")
	fmt.Printf("%+v\n", res)
	fmt.Println("=============")

	tx.Commit()
	return nil

}
