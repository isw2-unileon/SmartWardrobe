package services

import (
	"backend/database"
	"backend/models"
)

func GetUserGarments(userID string) ([]models.Garment, error) {
	rows, err := database.DB.Query(`
		SELECT id, user_id, type, color, style, image_url
		FROM garments
		WHERE user_id = $1
	`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var garments []models.Garment

	for rows.Next() {
		var g models.Garment
		err := rows.Scan(&g.ID, &g.UserID, &g.Type, &g.Color, &g.Style, &g.ImageURL)
		if err != nil {
			return nil, err
		}
		garments = append(garments, g)
	}

	return garments, nil
}

func CreateGarment(g models.Garment) error {
	_, err := database.DB.Exec(`
		INSERT INTO garments (user_id, type, color, style, image_url)
		VALUES ($1, $2, $3, $4, $5)
	`,
		g.UserID,
		g.Type,
		g.Color,
		g.Style,
		g.ImageURL,
	)

	return err
}