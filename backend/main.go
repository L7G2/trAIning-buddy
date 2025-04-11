package main

import (
	"backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db := InitDB()
	defer db.Close()

	RunMigrations(db)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "DB + backend status = OK"})
	})
	routes.RegisterRoutes(router, db)
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
