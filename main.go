package main

import (
	config "golang-mvc-rest-api/config"

	"log"

	"github.com/gorilla/mux"
)

func main() {
	dbConn, err := config.GetPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	router := mux.NewRouter()

}
