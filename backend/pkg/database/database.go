package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var db *sql.DB

func Connect() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("[SQL] Error connecting to database: %v", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("[SQL] Error pinging database: %v", err)
	}
	log.Println("[SQL] Connected to database")
	return db, nil
}

func Close() {
	if db != nil {
		_ = db.Close()
		log.Println("[SQL] Closed database connection")
	}
}

func RunMigrations(db *sql.DB) error {
	migrationFile := "migrations/001_init.sql"
	connect, err := os.ReadFile(migrationFile)
	if err != nil {
		return fmt.Errorf("[SQL] Error reading migration file: %v", err)
	}
	_, err = db.Exec(string(connect))
	if err != nil {
		return fmt.Errorf("[SQL] Error executing migration file: %v", err)
	}
	log.Println("[SQL] Migrations successfully executed")
	return nil
}
