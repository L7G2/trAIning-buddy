package models

import "time"

type DietPlan struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Meals     []Meal    `json:"meals,omitempty"`
}
