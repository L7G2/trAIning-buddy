package app

import (
	"backend/internal/middleware"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func SetupMiddleware(router *gin.Engine) {
	router.Use(middleware.CORSMiddleware())
}

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Backend is up and running"})
	})
	auth := handlers.NewAuthHandler(db)
	router.POST("/register", auth.Register)
	router.POST("/login", auth.Login)

	router.GET("/me", middleware.AuthMiddleware(), auth.Me)

}
