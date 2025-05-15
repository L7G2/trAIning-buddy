package repositories

import (
	"backend/internal/domain/models"
	"database/sql"
	"fmt"
)

type ProfileRepository struct {
	DB *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{DB: db}
}

func (r *ProfileRepository) GetByUserID(userID int) (*models.Profile, error) {
	row := r.DB.QueryRow(`SELECT user_id, age, height, weight, gender, goal FROM profiles WHERE user_id = $1`, userID)
	var p models.Profile
	err := row.Scan(&p.UserID, &p.Age, &p.Height, &p.Weight, &p.Gender, &p.Goal)
	if err != nil {
		fmt.Println("SCAN ERROR: ", err)
		return nil, err
	}
	return &p, nil
}

func (r *ProfileRepository) CreateOrUpdate(p *models.Profile) error {
	_, err := r.DB.Exec(`
		INSERT INTO profiles (user_id, age, height, weight, gender, goal)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (user_id)
		DO UPDATE SET age = $2, height = $3, weight = $4, gender = $5, goal = $6
	`, p.UserID, p.Age, p.Height, p.Weight, p.Gender, p.Goal)
	return err
}
