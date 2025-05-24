package repositories

import (
	"backend/internal/domain/models"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(u *models.User) error {
	query := `
		INSERT INTO users (username, password, role)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	return r.DB.QueryRow(query, u.Username, u.Password, u.Role).Scan(&u.ID)
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	rows, err := r.DB.Query(`SELECT id, username, role FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Role); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepository) GetByID(id int) (models.User, error) {
	var u models.User
	query := `SELECT id, username, role FROM users WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&u.ID, &u.Username, &u.Role)
	return u, err
}
