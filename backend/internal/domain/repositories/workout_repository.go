package repositories

import (
	"backend/internal/domain/models"
	"database/sql"
)

type WorkoutRepository struct {
	DB *sql.DB
}

func NewWorkoutRepository(db *sql.DB) *WorkoutRepository {
	return &WorkoutRepository{DB: db}
}

func (r *WorkoutRepository) Create(w *models.Workout) error {
	query := `INSERT INTO workouts (plan_id, workout_date, notes) VALUES ($1, $2, $3) RETURNING id`
	return r.DB.QueryRow(query, w.PlanID, w.WorkoutDate, w.Notes).Scan(&w.ID)
}

func (r *WorkoutRepository) GetByPlanID(planID int) ([]models.Workout, error) {
	rows, err := r.DB.Query(`SELECT id, plan_id, workout_date, notes FROM workouts WHERE plan_id = $1`, planID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workouts []models.Workout
	for rows.Next() {
		var w models.Workout
		if err := rows.Scan(&w.ID, &w.PlanID, &w.WorkoutDate, &w.Notes); err != nil {
			return nil, err
		}
		workouts = append(workouts, w)
	}
	return workouts, nil
}
