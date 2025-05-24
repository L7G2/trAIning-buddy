package repositories

import (
	"backend/internal/domain/models"
	"database/sql"
)

type ProgressReportRepository struct {
	DB *sql.DB
}

func NewProgressReportRepository(db *sql.DB) *ProgressReportRepository {
	return &ProgressReportRepository{DB: db}
}

func (r *ProgressReportRepository) Create(report *models.ProgressReport) error {
	query := `
		INSERT INTO progress_reports (user_id, date, weight, notes)
		VALUES ($1, $2, $3, $4) RETURNING id
	`
	return r.DB.QueryRow(
		query,
		report.UserID,
		report.Date,
		report.Weight,
		report.Notes,
	).Scan(&report.ID)
}

func (r *ProgressReportRepository) GetByUserID(userID int) ([]models.ProgressReport, error) {
	rows, err := r.DB.Query(`
		SELECT id, user_id, date, weight, notes
		FROM progress_reports
		WHERE user_id = $1
		ORDER BY date DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []models.ProgressReport
	for rows.Next() {
		var rpt models.ProgressReport
		if err := rows.Scan(&rpt.ID, &rpt.UserID, &rpt.Date, &rpt.Weight, &rpt.Notes); err != nil {
			return nil, err
		}
		reports = append(reports, rpt)
	}
	return reports, nil
}
