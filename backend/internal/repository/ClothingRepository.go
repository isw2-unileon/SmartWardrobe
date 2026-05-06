package repository

import (
	"backend/database"
	"backend/internal/models"
)

// Get valid clothing items (JOIN + filtro por temperatura)
func GetValidClothingItems(userID string, temp float64) ([]models.ClothingItem, error) {
	rows, err := database.DB.Query(`
		SELECT 
			ci.id,
			ci.type_id,
			ci.color_id,
			ci.style_id,
			ci.image_url,
			ci.user_id,
			mt.category_id,
			mt.warmth,
			mt.layer
		FROM clothing_items ci
		JOIN master_types mt ON ci.type_id = mt.id
		WHERE ci.user_id = $1
		AND $2 BETWEEN mt.min_temp AND mt.max_temp
	`, userID, temp)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.ClothingItem

	for rows.Next() {
		var i models.ClothingItem
		err := rows.Scan(
			&i.ID,
			&i.TypeID,
			&i.ColorID,
			&i.StyleID,
			&i.ImageURL,
			&i.UserID,
			&i.CategoryID,
			&i.Warmth,
			&i.Layer,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	return items, nil
}

// Get weather rule
func GetWeatherRule(temp float64) (*models.WeatherRule, error) {
	row := database.DB.QueryRow(`
		SELECT id, min_temp, max_temp, required_warmth, max_upper_layers
		FROM weather_rules
		WHERE $1 BETWEEN min_temp AND max_temp
		LIMIT 1
	`, temp)

	var rule models.WeatherRule

	err := row.Scan(
		&rule.ID,
		&rule.MinTemp,
		&rule.MaxTemp,
		&rule.RequiredWarmth,
		&rule.MaxUpperLayers,
	)

	if err != nil {
		return nil, err
	}

	return &rule, nil
}