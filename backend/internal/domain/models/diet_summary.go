package models

type DietSummary struct {
	DietPlanID    int `json:"diet_plan_id"`
	TotalCalories int `json:"total_calories"`
	TotalProteins int `json:"total_proteins"`
	TotalFats     int `json:"total_fats"`
	TotalCarbs    int `json:"total_carbs"`
}
