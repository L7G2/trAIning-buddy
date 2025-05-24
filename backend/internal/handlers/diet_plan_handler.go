package handlers

import (
	"backend/internal/domain/models"
	"backend/internal/domain/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DietPlanHandler struct {
	Repo *repositories.DietPlanRepository
}

func NewDietPlanHandler(repo *repositories.DietPlanRepository) *DietPlanHandler {
	return &DietPlanHandler{Repo: repo}
}

func (h *DietPlanHandler) Create(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	plan := models.DietPlan{
		UserID: userID,
	}

	if err := h.Repo.Create(&plan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create failed"})
		return
	}
	c.JSON(http.StatusCreated, plan)
}

func (h *DietPlanHandler) List(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	plans, err := h.Repo.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Query failed"})
		return
	}
	c.JSON(http.StatusOK, plans)
}

func (h *DietPlanHandler) Summary(c *gin.Context) {
	planID, err := strconv.Atoi(c.Param("dietPlanID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid diet plan ID"})
		return
	}

	summary, err := h.Repo.GetSummary(planID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Summary failed"})
		return
	}

	c.JSON(http.StatusOK, summary)
}
