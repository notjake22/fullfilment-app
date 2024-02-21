package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	DBConn *sql.DB
)

func InitDbConnection() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_PRIVATE_URL")+"?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	DBConn = db
}
