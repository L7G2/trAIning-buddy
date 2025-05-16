package handlers

import (
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ExerciseHandler struct {
	Repo *repositories.ExerciseRepository
}

func NewExerciseHandler(repo *repositories.ExerciseRepository) *ExerciseHandler {
	return &ExerciseHandler{Repo: repo}
}

func (h *ExerciseHandler) Create(c *gin.Context) {
	workoutID, _ := strconv.Atoi(c.Param("workoutID"))

	var e models.Exercise
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	e.WorkoutID = workoutID

	if err := h.Repo.Create(&e); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create failed"})
		return
	}
	c.JSON(http.StatusCreated, e)
}

func (h *ExerciseHandler) List(c *gin.Context) {
	workoutID, _ := strconv.Atoi(c.Param("workoutID"))

	exercises, err := h.Repo.GetByWorkoutID(workoutID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Query failed"})
		return
	}
	c.JSON(http.StatusOK, exercises)
}
