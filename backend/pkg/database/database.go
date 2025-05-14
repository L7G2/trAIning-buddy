package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func Connect() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("[SQL] Error connecting: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("[SQL] Error pinging DB: %w", err)
	}

	log.Println("[SQL] Connected to database")
	return db, nil
}

func RunMigrations(db *sql.DB) error {
	migrationFiles := []string{
		"migrations/001_init.sql",
		"migrations/002_products.sql",
	}

	for _, file := range migrationFiles {
		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("[SQL] Error reading migration %s: %w", file, err)
		}

		if _, err := db.Exec(string(content)); err != nil {
			return fmt.Errorf("[SQL] Error executing %s: %w", file, err)
		}

		log.Printf("[SQL] Migration %s applied", file)
	}

	return nil
}
