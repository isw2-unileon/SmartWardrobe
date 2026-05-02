package repository

import (
	"backend/internal/models"
	"database/sql"
)

// MasterTypeRepository handles database operations for the table master_types
type MasterTypeRepository struct {
	db *sql.DB
}

func NewMasterTypeRepository(db *sql.DB) *MasterTypeRepository {
	return &MasterTypeRepository{db: db}
}

func (r *MasterTypeRepository) GetAll() ([]models.MasterType, error) {
	var types []models.MasterType
	query := "SELECT * FROM master_types"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	// The rows are closed to release connection
	defer rows.Close()

	// Itemized for each returned record
	for rows.Next() {
		var typeM models.MasterType

		if err := rows.Scan(&typeM.ID, &typeM.Name, &typeM.CategoryId, &typeM.MinTemp, &typeM.MaxTemp, &typeM.Warmth, &typeM.Layer); err != nil {
			return nil, err
		}
		types = append(types, typeM)
	}

	// Checked if there was an error during the iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return types, err
}
