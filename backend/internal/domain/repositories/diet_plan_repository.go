package repositories

import (
	"backend/internal/domain/models"
	"database/sql"
)

type DietPlanRepository struct {
	DB *sql.DB
}

func NewDietPlanRepository(db *sql.DB) *DietPlanRepository {
	return &DietPlanRepository{DB: db}
}

func (r *DietPlanRepository) Create(plan *models.DietPlan) error {
	query := `
		INSERT INTO diet_plans (user_id)
		VALUES ($1)
		RETURNING id, created_at
	`
	return r.DB.QueryRow(query, plan.UserID).Scan(&plan.ID, &plan.CreatedAt)
}

func (r *DietPlanRepository) GetByUserID(userID int) ([]models.DietPlan, error) {
	rows, err := r.DB.Query(`
		SELECT id, user_id, created_at
		FROM diet_plans
		WHERE user_id = $1
		ORDER BY created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plans []models.DietPlan
	for rows.Next() {
		var plan models.DietPlan
		if err := rows.Scan(&plan.ID, &plan.UserID, &plan.CreatedAt); err != nil {
			return nil, err
		}

		meals, err := r.getMealsForPlan(plan.ID)
		if err != nil {
			return nil, err
		}
		plan.Meals = meals

		plans = append(plans, plan)
	}
	return plans, nil
}

func (r *DietPlanRepository) getMealsForPlan(planID int) ([]models.Meal, error) {
	rows, err := r.DB.Query(`
		SELECT id, diet_plan_id, meal_order, name, calories, proteins, fats, carbs
		FROM meals
		WHERE diet_plan_id = $1
		ORDER BY meal_order
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

func (r *DietPlanRepository) GetSummary(planID int) (models.DietSummary, error) {
	var summary models.DietSummary
	query := `
		SELECT
			diet_plan_id,
			COALESCE(SUM(calories), 0),
			COALESCE(SUM(proteins), 0),
			COALESCE(SUM(fats), 0),
			COALESCE(SUM(carbs), 0)
		FROM meals
		WHERE diet_plan_id = $1
		GROUP BY diet_plan_id
	`
	row := r.DB.QueryRow(query, planID)
	err := row.Scan(&summary.DietPlanID, &summary.TotalCalories, &summary.TotalProteins, &summary.TotalFats, &summary.TotalCarbs)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.DietSummary{DietPlanID: planID}, nil
		}
		return summary, err
	}
	return summary, nil
}
