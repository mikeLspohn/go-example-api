package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//DBConnection opens a new db connection and returns the db to be used
//Make sure to close the connection
func DBConnection() *sqlx.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://michael@localhost/gotest?sslmode=disable"
	}

	db, err := sqlx.Open("postgres", dbURL)

	// Fail before returning db so it doesn't need to be checked elsewhere
	if err != nil {
		log.Fatal("Error: Postgres connection error")
	}

	return db
}
