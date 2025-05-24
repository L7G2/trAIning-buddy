package handlers

import (
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MealHandler struct {
	Repo *repositories.MealRepository
}

func NewMealHandler(repo *repositories.MealRepository) *MealHandler {
	return &MealHandler{Repo: repo}
}

func (h *MealHandler) Create(c *gin.Context) {
	planID, err := strconv.Atoi(c.Param("dietPlanID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid diet plan ID"})
		return
	}

	var meal models.Meal
	if err := c.ShouldBindJSON(&meal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	meal.DietPlanID = planID

	if err := h.Repo.Create(&meal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create failed"})
		return
	}
	c.JSON(http.StatusCreated, meal)
}

func (h *MealHandler) List(c *gin.Context) {
	planID, err := strconv.Atoi(c.Param("dietPlanID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid diet plan ID"})
		return
	}

	meals, err := h.Repo.GetByDietPlanID(planID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Query failed"})
		return
	}
	c.JSON(http.StatusOK, meals)
}
