package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/JulzDiverse/feedelphia/api"
	feedb "github.com/JulzDiverse/feedelphia/db"
)

const defaultPort = "8080"

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = defaultPort
	}

	dbURL := os.Getenv("DB_CONNECTION_STRING")
	if len(dbURL) == 0 {
		panic(errors.New("DB url not provided"))
	}

	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS photos(
title varchar(100) NOT NULL,
author varchar(100) NOT NULL,
hero varchar(50) NOT NULL,
data MEDIUMBLOB NOT NULL,
timestamp TIMESTAMP NOT NULL
);`)
	if err != nil {
		panic(err)
	}

	photobase := feedb.NewSQLPhotobase(db)
	handler := api.New(&photobase)

	http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
