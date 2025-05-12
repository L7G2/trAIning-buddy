package routes

import (
	"backend/middleware"
	"backend/models"
	"backend/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	router.POST("/register", func(c *gin.Context) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Role     string `json:"role"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		_, err = db.Exec("INSERT INTO users (username, password, role) VALUES ($1, $2, $3)", req.Username, hashedPassword, req.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "insert error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "registered successfully"})
	})
	router.POST("/login", func(c *gin.Context) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
			// Tymczasowe logowanie "admin" bez bazy
	if req.Username == "admin" && req.Password == "admin" {
		token, err := utils.GenerateTokenJWT("0", "trener")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "logged in as test admin",
			"token":   token,
			"user": gin.H{
				"id":       0,
				"username": "admin",
				"role":     "trener",
			},
		})
		return
	}
		var user models.User
		err := db.QueryRow("SELECT id, username, password, role FROM users WHERE username = $1", req.Username).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username/password"})
			return
		}

		if !utils.CheckPasswordHash(req.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username/password"})
			return
		}
		token, err := utils.GenerateTokenJWT(strconv.Itoa(user.ID), user.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "logged in",
			"token":   token,
			"user": gin.H{
				"id":       user.ID,
				"username": user.Username,
				"role":     user.Role,
			},
		})
	})
	router.GET("/me", middleware.AuthMiddleware(), func(c *gin.Context) {
		userID := c.MustGet("userID").(int)
		role := c.MustGet("role").(string)
		c.JSON(http.StatusOK, gin.H{
			"userID": userID,
			"role":   role,
		})
	})
}
