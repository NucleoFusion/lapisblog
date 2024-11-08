package db

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectToDB() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", os.Getenv("PSQL_URI"))
	if err != nil {
		return db, err
	}

	return db, nil
}