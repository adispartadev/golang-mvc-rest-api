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
	err := db.DB.QueryRow(query, user.Username).Scan(&userData.ID, &userData.Username, &userData.Password)

	if err == sql.ErrNoRows {
		return userData, errors.New("user is not found")
	}

	if err != nil {
		return userData, err
	}

	return userData, nil
}
