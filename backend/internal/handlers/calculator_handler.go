package handlers

import (
	"backend/internal/domain/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CalculatorHandler struct {
	DB *sql.DB
}

func NewCalculatorHandler(db *sql.DB) *CalculatorHandler {
	return &CalculatorHandler{DB: db}
}

type CalcItem struct {
	ProductID int     `json:"product_id"`
	Grams     float64 `json:"grams"`
}

type CalcRequest struct {
	Items []CalcItem `json:"items"`
}

type CalcResult struct {
	Calories float64 `json:"calories"`
	Protein  float64 `json:"protein"`
	Fat      float64 `json:"fat"`
	Carbs    float64 `json:"carbs"`
}

func (h *CalculatorHandler) Calculate(c *gin.Context) {
	var req CalcRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var result CalcResult

	for _, item := range req.Items {
		var product models.FoodProduct

		err := h.DB.QueryRow(`
		SELECT id, name, calories, protein, fat, carbs
		FROM food_products WHERE id = $1
	`, item.ProductID).Scan(
			&product.ID, &product.Name, &product.Calories,
			&product.Protein, &product.Fat, &product.Carbs,
		)

		if err != nil {
			continue
		}

		factor := item.Grams / 100.0
		result.Calories += product.Calories * factor
		result.Protein += product.Protein * factor
		result.Fat += product.Fat * factor
		result.Carbs += product.Carbs * factor
	}

	c.JSON(http.StatusOK, result)
}
