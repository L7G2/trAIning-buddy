package handlers

import (
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type WorkoutHandler struct {
	Repo *repositories.WorkoutRepository
}

func NewWorkoutHandler(repo *repositories.WorkoutRepository) *WorkoutHandler {
	return &WorkoutHandler{Repo: repo}
}

func (h *WorkoutHandler) Create(c *gin.Context) {
	planID, _ := strconv.Atoi(c.Param("planID"))

	var w models.Workout
	if err := c.ShouldBindJSON(&w); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	w.PlanID = planID

	if err := h.Repo.Create(&w); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create failed"})
		return
	}
	c.JSON(http.StatusCreated, w)
}

func (h *WorkoutHandler) List(c *gin.Context) {
	planID, _ := strconv.Atoi(c.Param("planID"))

	workouts, err := h.Repo.GetByPlanID(planID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Query failed"})
		return
	}
	c.JSON(http.StatusOK, workouts)
}
