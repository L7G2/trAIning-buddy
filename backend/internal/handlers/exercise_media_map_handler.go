package handlers

import (
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type ExerciseMediaHandler struct {
	Repo *repositories.ExerciseMediaRepository
}

func NewExerciseMediaHandler(repo *repositories.ExerciseMediaRepository) *ExerciseMediaHandler {
	return &ExerciseMediaHandler{Repo: repo}
}

func (h *ExerciseMediaHandler) CreateLink(c *gin.Context) {
	var link models.ExerciseMediaMap
	if err := c.ShouldBindJSON(&link); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Repo.Create(&link); err != nil {
		log.Println("Create error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to link media to exercise"})
		return
	}
	c.JSON(http.StatusCreated, link)
}

func (h *ExerciseMediaHandler) GetMediaByExerciseID(c *gin.Context) {
	exerciseID, err := strconv.Atoi(c.Param("exerciseID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exercise ID"})
		return
	}

	files, err := h.Repo.GetMediaFilesByExerciseID(exerciseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get media"})
		return
	}
	c.JSON(http.StatusOK, files)
}
