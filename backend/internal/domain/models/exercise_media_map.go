package models

type ExerciseMediaMap struct {
	ID         int `json:"id"`
	ExerciseID int `json:"exercise_id"`
	MediaID    int `json:"media_id"`
}
