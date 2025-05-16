package main

import (
	"backend/internal/app"
	"backend/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("[WARN] Error loading .env file")
	}
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("[DB] Connection error: %v", err)
	}
	defer db.Close()

	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("[DB] Migrations error: %v", err)
	}

	router := gin.Default()
	app.SetupMiddleware(router)
	app.SetupRoutes(router, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("[APP] Server listening on " + port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("[APP] Server failed: %v", err)
	}
}
