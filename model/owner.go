package model

import (
	"database/sql"
	"errors"
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

	_, err = tx.Exec(query, owner.Name)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil

}

func DeleteOwner(ownerID int) error {
	const query = `DELETE FROM owners WHERE id = $1`
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	res, err := tx.Exec(query, ownerID)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return errors.New("data is not found")
	}

	if rowsAffected > 1 {
		return errors.New("Strange behaviour. Total affected : " + string(rowsAffected))
	}

	tx.Commit()
	return nil

}

func UpdateOwner(owner *e.Owner) error {

	const query = `UPDATE owners SET name = $2 WHERE id = $1`
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	res, err := tx.Exec(query, owner.ID, owner.Name)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return errors.New("data is not found")
	}

	if rowsAffected > 1 {
		tx.Rollback()
		return errors.New("Strange behaviour. Total affected is : " + string(rowsAffected))
	}

	tx.Commit()

	return nil

}
