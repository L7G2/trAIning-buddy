package repositories

import (
	"backend/internal/domain/models"
	"database/sql"
)

type MealRepository struct {
	DB *sql.DB
}

func NewMealRepository(db *sql.DB) *MealRepository {
	return &MealRepository{DB: db}
}

func (r *MealRepository) Create(meal *models.Meal) error {
	query := `
		INSERT INTO meals (diet_plan_id, meal_order, name, calories, proteins, fats, carbs)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`
	return r.DB.QueryRow(query, meal.DietPlanID, meal.MealOrder, meal.Name, meal.Calories, meal.Proteins, meal.Fats, meal.Carbs).Scan(&meal.ID)
}

func (r *MealRepository) GetByDietPlanID(planID int) ([]models.Meal, error) {
	rows, err := r.DB.Query(`
		SELECT id, diet_plan_id, meal_order, name, calories, proteins, fats, carbs
		FROM meals
		WHERE diet_plan_id = $1
		ORDER BY meal_order ASC
	`, planID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meals []models.Meal
	for rows.Next() {
		var m models.Meal
		if err := rows.Scan(&m.ID, &m.DietPlanID, &m.MealOrder, &m.Name, &m.Calories, &m.Proteins, &m.Fats, &m.Carbs); err != nil {
			return nil, err
		}
		meals = append(meals, m)
	}
	return meals, nil
}
