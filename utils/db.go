package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func ConnectDB() *sqlx.DB {
	// get db env variable
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// create postgres db connection
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Panic("error saat koneksi db: ", err)
	}
	log.Println("berhasil terhubung dengan database")

	return db
}
