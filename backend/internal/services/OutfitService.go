package services

import (
    "backend/internal/models"
    "math/rand"
)

func getHigherWarmthItem(items []models.ClothingItem) *models.ClothingItem {
    if len(items) == 0 {
        return nil
    }

    maxWarmth := items[0].Warmth
    for _, item := range items {
        if item.Warmth > maxWarmth {
            maxWarmth = item.Warmth
        }
    }

    var candidates []int
    for i, item := range items {
        if item.Warmth == maxWarmth {
            candidates = append(candidates, i)
        }
    }

    randIndex := candidates[rand.Intn(len(candidates))]
    return &items[randIndex]
}

func GenerateOutfit(items []models.ClothingItem, rule *models.WeatherRule) *models.OutfitItem {

    // FIX: variable names corrected
    var upperwear, bottomwear, footwear []models.ClothingItem

    for _, item := range items {
        switch item.CategoryID {
        case 1:
            upperwear = append(upperwear, item)
        case 2:
            bottomwear = append(bottomwear, item)
        case 3:
            footwear = append(footwear, item)
        }
    }

    outfit := &models.OutfitItem{
        Upperwear:  []*models.ClothingItem{},
        Bottomwear: nil,
        Footwear:   nil,
    }

    if len(bottomwear) > 0 {
        outfit.Bottomwear = getHigherWarmthItem(bottomwear)
    }

    if len(footwear) > 0 {
        outfit.Footwear = getHigherWarmthItem(footwear)
    }

    var warmth int64 = 0

    for layer := int64(0); layer < rule.MaxUpperLayers; layer++ {

        var layerItems []models.ClothingItem
        for _, item := range upperwear {
            if item.Layer != nil && *item.Layer == layer {
                layerItems = append(layerItems, item)
            }
        }

        chosen := getHigherWarmthItem(layerItems)
        if chosen != nil {
            outfit.Upperwear = append(outfit.Upperwear, chosen)
            warmth += chosen.Warmth

            if warmth >= rule.RequiredWarmth {
                break
            }
        }
    }

    return outfit
}
