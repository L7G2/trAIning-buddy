package handlers

import (
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProfileHandler struct {
	Repo *repositories.ProfileRepository
}

func NewProfileHandler(repo *repositories.ProfileRepository) *ProfileHandler {
	return &ProfileHandler{Repo: repo}
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(int)
	profile, err := h.Repo.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"profile": profile})
}
func (h *ProfileHandler) SaveProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(int)

	var input models.Profile
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.UserID = userID

	err := h.Repo.CreateOrUpdate(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profile saved"})
}
