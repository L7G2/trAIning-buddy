package handlers

import (
	"backend/internal/domain/models"
	"backend/pkg/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AuthHandler struct {
	DB *sql.DB
}

func NewAuthHandler(DB *sql.DB) *AuthHandler {
	return &AuthHandler{DB}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Role     string `json:"role" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	_, err = h.DB.Exec(`INSERT INTO users (username, password, role) VALUES ($1, $2, $3)`,
		req.Username, hashedPassword, req.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Successfully registered user"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	row := h.DB.QueryRow(`SELECT id, username, password, role FROM users WHERE username = $1`,
		req.Username)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid query data"})
		return
	}
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	token, err := utils.GenerateTokenJWT(strconv.Itoa(user.ID), user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "jwt failed"})
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
}

func (h *AuthHandler) Me(c *gin.Context) {
	userID := c.MustGet("user_id").(int)
	role := c.MustGet("role").(string)
	c.JSON(http.StatusOK, gin.H{
		"userID": userID,
		"role":   role,
	})
}
