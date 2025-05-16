package models

type TrainingPlan struct {
	ID          int    `json:"id"`
	ClientID    int    `json:"client_id"`
	TrainerID   int    `json:"trainer_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}
