package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func InitDB() *sql.DB {
	dbUser := "postgres"
	dbPass := "database"
	dbName := "tbdb"
	dbHost := "localhost"
	dbPort := "5432"

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("[DB] connection error %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("[DB] ping error %v", err)
	}

	log.Println("Successfully connected to database")
	return db
}

func RunMigrations(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		password TEXT NOT NULL,
		role VARCHAR(20) NOT NULL CHECK (role IN ('trener', 'uczen'))
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("[DB] migration error: %v", err)
	}

	log.Println("[DB] migration succeeded")
}
