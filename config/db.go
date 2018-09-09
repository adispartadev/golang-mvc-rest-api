package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// GetPostgresDB - get postgres db config connection
func GetPostgresDB() (*sql.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DB_NAME")

	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, databaseName)

	db, err := createDBConnection(desc)

	if err != nil {
		return nil, err
	}

	return db, nil

}

func createDBConnection(desc string) (*sql.DB, error) {
	db, err := sql.Open("postgres", desc)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db, nil
}
