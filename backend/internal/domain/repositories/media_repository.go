package repositories

import (
	"backend/internal/domain/models"
	"database/sql"
)

type MediaRepository struct {
	DB *sql.DB
}

func NewMediaRepository(db *sql.DB) *MediaRepository {
	return &MediaRepository{DB: db}
}

func (r *MediaRepository) Create(m *models.MediaFile) error {
	query := `
		INSERT INTO media_files (title, url, uploaded_by)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`
	return r.DB.QueryRow(query, m.Title, m.URL, m.UploadedBy).
		Scan(&m.ID, &m.CreatedAt)
}

func (r *MediaRepository) GetAll() ([]models.MediaFile, error) {
	rows, err := r.DB.Query(`SELECT id, title, url, uploaded_by, created_at FROM media_files`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []models.MediaFile
	for rows.Next() {
		var m models.MediaFile
		err := rows.Scan(&m.ID, &m.Title, &m.URL, &m.UploadedBy, &m.CreatedAt)
		if err != nil {
			return nil, err
		}
		files = append(files, m)
	}
	return files, nil
}

func (r *MediaRepository) GetByExerciseID(exerciseID int) ([]models.MediaFile, error) {
	rows, err := r.DB.Query(`
		SELECT mf.id, mf.title, mf.url, mf.uploaded_by, mf.created_at
		FROM media_files mf
		JOIN exercise_media_files emf ON mf.id = emf.media_file_id
		WHERE emf.exercise_id = $1
	`, exerciseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []models.MediaFile
	for rows.Next() {
		var m models.MediaFile
		if err := rows.Scan(&m.ID, &m.Title, &m.URL, &m.UploadedBy, &m.CreatedAt); err != nil {
			return nil, err
		}
		files = append(files, m)
	}
	return files, nil
}
