package main

import (
	"backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := InitDB()
	defer db.Close()

	RunMigrations(db)
	router := gin.Default()

		router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "DB + backend status = OK"})
	})
	routes.RegisterRoutes(router, db)
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
