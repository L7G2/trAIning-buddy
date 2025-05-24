package handlers

import (
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MediaHandler struct {
	Repo *repositories.MediaRepository
}

func NewMediaHandler(repo *repositories.MediaRepository) *MediaHandler {
	return &MediaHandler{Repo: repo}
}

func (h *MediaHandler) Create(c *gin.Context) {
	var m models.MediaFile
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userID, exists := c.Get("userID"); exists {
		id := userID.(int)
		m.UploadedBy = &id
	}

	if err := h.Repo.Create(&m); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create failed"})
		return
	}

	c.JSON(http.StatusCreated, m)
}

func (h *MediaHandler) List(c *gin.Context) {
	files, err := h.Repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Query failed"})
		return
	}
	c.JSON(http.StatusOK, files)
}

func (h *MediaHandler) GetByExerciseID(c *gin.Context) {
	exerciseID, err := strconv.Atoi(c.Param("exerciseID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exercise ID"})
		return
	}

	files, err := h.Repo.GetByExerciseID(exerciseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Query failed"})
		return
	}
	c.JSON(http.StatusOK, files)
}
