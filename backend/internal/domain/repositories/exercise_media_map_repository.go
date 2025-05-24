package repositories

import (
	"backend/internal/domain/models"
	"database/sql"
)

type ExerciseMediaRepository struct {
	DB *sql.DB
}

func NewExerciseMediaRepository(db *sql.DB) *ExerciseMediaRepository {
	return &ExerciseMediaRepository{DB: db}
}

func (r *ExerciseMediaRepository) Create(link *models.ExerciseMediaMap) error {
	query := `INSERT INTO exercise_media_map (exercise_id, media_id) VALUES ($1, $2) RETURNING id`
	return r.DB.QueryRow(query, link.ExerciseID, link.MediaID).Scan(&link.ID)
}

func (r *ExerciseMediaRepository) GetMediaFilesByExerciseID(exerciseID int) ([]models.MediaFile, error) {
	query := `
		SELECT m.id, m.title, m.url, m.uploaded_by, m.created_at
		FROM media_files m
		JOIN exercise_media_map emf ON m.id = emf.media_id
		WHERE emf.exercise_id = $1
	`

	rows, err := r.DB.Query(query, exerciseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mediaFiles []models.MediaFile
	for rows.Next() {
		var m models.MediaFile
		err := rows.Scan(&m.ID, &m.Title, &m.URL, &m.UploadedBy, &m.CreatedAt)
		if err != nil {
			return nil, err
		}
		mediaFiles = append(mediaFiles, m)
	}
	return mediaFiles, nil
}
