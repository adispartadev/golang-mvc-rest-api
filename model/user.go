package model

import (
	"database/sql"
	"errors"
	"golang-mvc-rest-api/db"
	e "golang-mvc-rest-api/entity"
)

func InsertUser(user *e.User) error {
	const query = `INSERT INTO users (username, password) VALUES ($1, $2)`
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, user.Username, user.Password)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func GetUserByUsername(user *e.User) (e.User, error) {
	var userData e.User
	const query = `SELECT * FROM users WHERE username = $1`
	err := db.DB.QueryRow(query, user.Username).Scan(&userData.ID, &userData.Username, &userData.Password, &userData.RefreshToken)

	if err == sql.ErrNoRows {
		return userData, errors.New("user is not found")
	}

	if err != nil {
		return userData, err
	}

	return userData, nil
}

func GetUserByRefreshToken(user *e.User, refreshToken string) error {
	var countUser int

	const query = `SELECT count(*) FROM users WHERE username = $1 and refresh_token = $2`
	err := db.DB.QueryRow(query, user.Username, refreshToken).Scan(&countUser)

	if err == sql.ErrNoRows {
		return err
	}

	if err != nil {
		return err
	}

	if countUser > 1 {
		return errors.New("Strange behaviour: Got more than 1 user")
	}

	return nil
}

func UpdateRefreshToken(user *e.User, refreshToken string) error {
	const query = `UPDATE users SET refresh_token = $2 WHERE username = $1`
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	res, err := tx.Exec(query, user.Username, refreshToken)
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
		return errors.New("Strange behaviour. Total affected is :" + string(rowsAffected))
	}

	tx.Commit()
	return nil
}
