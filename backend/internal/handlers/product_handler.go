package handlers

import (
	"backend/internal/domain/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductHandler struct {
	DB *sql.DB
}

func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{DB: db}
}

func (h *ProductHandler) List(c *gin.Context) {
	search := c.Query("search")
	query := "SELECT id, name, calories, protein, fat, carbs FROM food_products"
	args := []interface{}{}

	if search != "" {
		query += " WHERE LOWER(name) LIKE LOWER($1)"
		args = append(args, "%"+search+"%")
	}

	query += " ORDER BY name"

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var products []models.FoodProduct

	for rows.Next() {
		var p models.FoodProduct
		if err := rows.Scan(&p.ID, &p.Name, &p.Calories, &p.Protein, &p.Fat, &p.Carbs); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Scan error"})
			return
		}
		products = append(products, p)
	}

	c.JSON(http.StatusOK, products)
}
