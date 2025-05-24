package models

type Meal struct {
	ID         int    `json:"id"`
	DietPlanID int    `json:"diet_plan_id"`
	MealOrder  int    `json:"meal_order"`
	Name       string `json:"name"`
	Calories   int    `json:"calories"`
	Proteins   int    `json:"proteins"`
	Fats       int    `json:"fats"`
	Carbs      int    `json:"carbs"`
}
