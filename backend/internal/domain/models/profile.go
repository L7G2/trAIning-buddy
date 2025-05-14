package models

type Profile struct {
	UserID int    `json:"user_id"`
	Age    int    `json:"age"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Gender string `json:"gender"`
	Goal   string `json:"goal"`
}
