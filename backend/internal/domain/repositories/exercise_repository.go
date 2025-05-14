package repositories

import (
	"backend/internal/domain/models"
	"database/sql"
)

type ExerciseRepository struct {
	DB *sql.DB
}

func NewExerciseRepository(db *sql.DB) *ExerciseRepository {
	return &ExerciseRepository{DB: db}
}

func (r *ExerciseRepository) Create(e *models.Exercise) error {
	query := `
		INSERT INTO exercises (workout_id, name, sets, reps, weight, instructions)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
	`
	return r.DB.QueryRow(query, e.WorkoutID, e.Name, e.Sets, e.Reps, e.Weight, e.Instructions).Scan(&e.ID)
}

func (r *ExerciseRepository) GetByWorkoutID(workoutID int) ([]models.Exercise, error) {
	rows, err := r.DB.Query(`SELECT id, workout_id, name, sets, reps, weight, instructions FROM exercises WHERE workout_id = $1`, workoutID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercises []models.Exercise
	for rows.Next() {
		var e models.Exercise
		if err := rows.Scan(&e.ID, &e.WorkoutID, &e.Name, &e.Sets, &e.Reps, &e.Weight, &e.Instructions); err != nil {
			return nil, err
		}
		exercises = append(exercises, e)
	}
	return exercises, nil
}
