package models

type Exercise struct {
	ID           int    `json:"id"`
	WorkoutID    int    `json:"workout_id"`
	Name         string `json:"name"`
	Sets         int    `json:"sets"`
	Reps         int    `json:"reps"`
	Weight       int    `json:"weight"`
	Instructions string `json:"instructions"`
}
