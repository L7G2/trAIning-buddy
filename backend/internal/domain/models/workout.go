package models

type Workout struct {
	ID          int    `json:"id"`
	PlanID      int    `json:"plan_id"`
	WorkoutDate string `json:"workout_date"`
	Notes       string `json:"notes"`
}
