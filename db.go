package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DBConnection() *sqlx.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db_url := os.Getenv("DATABASE_URL")
	if db_url == "" {
		db_url = "postgres://michael@localhost/gotest?sslmode=disable"
	}

	db, err := sqlx.Open("postgres", db_url)

	// Fail before returning db so it doesn't need to be checked elsewhere
	if err != nil {
		log.Fatal("Error: Postgres connection error")
	}

	return db
}
