package main

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func migrate(db *sqlx.DB) {
	// create auth table migration
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS auth (
			id SERIAL PRIMARY KEY,
			email VARCHAR(100) NOT NULL,
			password VARCHAR(250) NOT NULL,
			role VARCHAR(10) CHECK (role IN ('merchant', 'admin'))
		)
	`)
	if err != nil {
		log.Panic("error saat migrasi tabel auth: ", err)
	}

	// create categories table migration
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL
		)
	`)
	if err != nil {
		log.Panic("error saat migrasi tabel categories: ", err)
	}

	// create products table migration
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			merchant_id INT NOT NULL REFERENCES auth (id),
			name VARCHAR(100) NOT NULL,
			stock SMALLINT NOT NULL,
			price INT NOT NULL,
			category_id INT NOT NULL REFERENCES categories (id)
		)
	`)
	if err != nil {
		log.Panic("error saat migrasi tabel products: ", err)
	}
}
