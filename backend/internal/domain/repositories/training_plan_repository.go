package repositories

import (
	"backend/internal/domain/models"
	"database/sql"
)

type TrainingPlanRepository struct {
	DB *sql.DB
}

func NewTrainingPlanRepository(db *sql.DB) *TrainingPlanRepository {
	return &TrainingPlanRepository{DB: db}
}

func (r *TrainingPlanRepository) Create(plan *models.TrainingPlan) error {
	query := `
		INSERT INTO training_plans (client_id, trainer_id, name, description, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5, NULLIF($6, '')) RETURNING id
	`
	return r.DB.QueryRow(query,
		plan.ClientID,
		plan.TrainerID,
		plan.Name,
		plan.Description,
		plan.StartDate,
		plan.EndDate,
	).Scan(&plan.ID)
}

func (r *TrainingPlanRepository) GetByClientID(clientID int) ([]models.TrainingPlan, error) {
	rows, err := r.DB.Query(`SELECT id, client_id, trainer_id, name, description, start_date, COALESCE(end_date, '') FROM training_plans WHERE client_id = $1`, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plans []models.TrainingPlan
	for rows.Next() {
		var p models.TrainingPlan
		err := rows.Scan(&p.ID, &p.ClientID, &p.TrainerID, &p.Name, &p.Description, &p.StartDate, &p.EndDate)
		if err != nil {
			return nil, err
		}
		plans = append(plans, p)
	}
	return plans, nil
}
