package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	dbUser := "postgres"
	dbPassword := "database"
	dbName := "tbdb"
	dbHost := "localhost"
	dbPort := "5432"

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("[DB] connection error %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("[DB] close error %v", err)
		}
	}(db)

	err = db.Ping()
	if err != nil {
		log.Fatalf("[DB] ping error %v", err)
	}
	log.Println("Successfully connected to database")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DB + backend status OK",
		})
	})

	err = router.Run(":8080")
	if err != nil {
		return
	}
}
