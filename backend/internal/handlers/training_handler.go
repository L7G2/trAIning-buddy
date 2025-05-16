package handlers

import (
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TrainingPlanHandler struct {
	Repo *repositories.TrainingPlanRepository
}

func NewTrainingPlanHandler(repo *repositories.TrainingPlanRepository) *TrainingPlanHandler {
	return &TrainingPlanHandler{Repo: repo}
}

func (h *TrainingPlanHandler) Create(c *gin.Context) {
	role := c.GetString("role")
	if role != "trener" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only trainers can create plans"})
		return
	}

	trainerID, _ := strconv.Atoi(c.GetString("userID"))

	var plan models.TrainingPlan
	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plan.TrainerID = trainerID

	if err := h.Repo.Create(&plan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create plan"})
		return
	}
	c.JSON(http.StatusCreated, plan)
}

func (h *TrainingPlanHandler) GetMyPlans(c *gin.Context) {
	userID, _ := strconv.Atoi(c.GetString("userID"))

	plans, err := h.Repo.GetByClientID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot load plans"})
		return
	}
	c.JSON(http.StatusOK, plans)
}
