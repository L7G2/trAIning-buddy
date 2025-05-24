package models

import "time"

type ProgressReport struct {
	ID     int  `json:"id"`
	UserID int  `json:"user_id"`
	User   User `json:"-"`

	Date   time.Time `json:"date"`
	Weight float64   `json:"weight"`
	Notes  string    `json:"notes"`
}
