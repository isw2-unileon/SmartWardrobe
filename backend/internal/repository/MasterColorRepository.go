package repository

import (
	"backend/internal/models"
	"database/sql"
)

// MasterColorRepository handles database operations for the table master_colors
type MasterColorRepository struct {
	db *sql.DB
}

func NewMasterColorRepository(db *sql.DB) *MasterColorRepository {
	return &MasterColorRepository{db: db}
}

func (r *MasterColorRepository) GetAll() ([]models.MasterColor, error) {
	var colors []models.MasterColor
	query := "SELECT * FROM master_colors"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	// The rows are closed to release connection
	defer rows.Close() 

    // Itemized for each returned record
    for rows.Next() {
        var color models.MasterColor
        
        if err := rows.Scan(&color.ID, &color.Name); err != nil {
            return nil, err
        }
        colors = append(colors, color)
    }

    // Checked if there was an error during the iteration
    if err := rows.Err(); err != nil {
        return nil, err
    }

	return colors, err
}
