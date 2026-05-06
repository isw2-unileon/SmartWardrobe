package services

import (
    "backend/internal/models"
    "testing"
)

// helper
func intPtr(i int64) *int64 { return &i }

// ------------------------------------------------------------
// 1. getHigherWarmthItem — basic functionality
// ------------------------------------------------------------
func testgetHigherWarmthItem(t *testing.T) {
    items := []models.ClothingItem{
        {ID: 1, Warmth: 2},
        {ID: 2, Warmth: 5},
        {ID: 3, Warmth: 5},
    }

    result := getHigherWarmthItem(items)

    if result == nil {
        t.Fatal("Expected item, got nil")
    }

    if result.Warmth != 5 {
        t.Errorf("Expected warmth 5, got %d", result.Warmth)
    }
}

// ------------------------------------------------------------
// 2. Normal case — all categories present
// ------------------------------------------------------------
func TestGenerateOutfit_Normal(t *testing.T) {

    items := []models.ClothingItem{
        // upperwear
        {ID: 1, CategoryID: 1, Warmth: 2, Layer: intPtr(0)},
        {ID: 2, CategoryID: 1, Warmth: 3, Layer: intPtr(1)},

        // bottomwear
        {ID: 3, CategoryID: 2, Warmth: 3},

        // footwear
        {ID: 4, CategoryID: 3, Warmth: 2},
    }

    rule := &models.WeatherRule{
        RequiredWarmth: 3,
        MaxUpperLayers: 3,
    }

    outfit := GenerateOutfit(items, rule)

    if outfit == nil {
        t.Fatal("Expected outfit, got nil")
    }

    if len(outfit.Upperwear) == 0 {
        t.Error("Expected at least one upperwear layer")
    }

    if outfit.Bottomwear == nil {
        t.Error("Expected bottomwear")
    }

    if outfit.Footwear == nil {
        t.Error("Expected footwear")
    }
}

// ------------------------------------------------------------
// 3. No upperwear
// ------------------------------------------------------------
func TestGenerateOutfit_NoUpperwear(t *testing.T) {

    items := []models.ClothingItem{
        {ID: 1, CategoryID: 2, Warmth: 3}, // bottom
        {ID: 2, CategoryID: 3, Warmth: 2}, // footwear
    }

    rule := &models.WeatherRule{
        RequiredWarmth: 5,
        MaxUpperLayers: 3,
    }

    outfit := GenerateOutfit(items, rule)

    if len(outfit.Upperwear) != 0 {
        t.Error("Expected no upperwear layers")
    }
}

// ------------------------------------------------------------
// 4. No bottomwear
// ------------------------------------------------------------
func TestGenerateOutfit_NoBottomwear(t *testing.T) {

    items := []models.ClothingItem{
        {ID: 1, CategoryID: 1, Warmth: 2, Layer: intPtr(0)}, // upper
        {ID: 2, CategoryID: 3, Warmth: 2},                   // footwear
    }

    rule := &models.WeatherRule{
        RequiredWarmth: 2,
        MaxUpperLayers: 3,
    }

    outfit := GenerateOutfit(items, rule)

    if outfit.Bottomwear != nil {
        t.Error("Expected nil bottomwear")
    }
}

// ------------------------------------------------------------
// 5. No footwear
// ------------------------------------------------------------
func TestGenerateOutfit_NoFootwear(t *testing.T) {

    items := []models.ClothingItem{
        {ID: 1, CategoryID: 1, Warmth: 2, Layer: intPtr(0)}, // upper
        {ID: 2, CategoryID: 2, Warmth: 3},                   // bottom
    }

    rule := &models.WeatherRule{
        RequiredWarmth: 2,
        MaxUpperLayers: 3,
    }

    outfit := GenerateOutfit(items, rule)

    if outfit.Footwear != nil {
        t.Error("Expected nil footwear")
    }
}

// ------------------------------------------------------------
// 6. Layer=nil should be ignored
// ------------------------------------------------------------
func TestGenerateOutfit_LayerNilIgnored(t *testing.T) {

    items := []models.ClothingItem{
        {ID: 1, CategoryID: 1, Warmth: 5, Layer: nil}, // should NOT be selected
        {ID: 2, CategoryID: 1, Warmth: 3, Layer: intPtr(0)},
    }

    rule := &models.WeatherRule{
        RequiredWarmth: 3,
        MaxUpperLayers: 3,
    }

    outfit := GenerateOutfit(items, rule)

    if len(outfit.Upperwear) != 1 {
        t.Errorf("Expected 1 upperwear layer, got %d", len(outfit.Upperwear))
    }

    if outfit.Upperwear[0].ID != 2 {
        t.Errorf("Expected ID 2, got %d", outfit.Upperwear[0].ID)
    }
}

// ------------------------------------------------------------
// 7. Warmth requirement stops early
// ------------------------------------------------------------
func TestGenerateOutfit_WarmthStopsEarly(t *testing.T) {

    items := []models.ClothingItem{
        {ID: 1, CategoryID: 1, Warmth: 5, Layer: intPtr(0)},
        {ID: 2, CategoryID: 1, Warmth: 5, Layer: intPtr(1)},
    }

    rule := &models.WeatherRule{
        RequiredWarmth: 5,
        MaxUpperLayers: 3,
    }

    outfit := GenerateOutfit(items, rule)

    if len(outfit.Upperwear) != 1 {
        t.Errorf("Expected only 1 layer, got %d", len(outfit.Upperwear))
    }
}

// ------------------------------------------------------------
// 8. MaxUpperLayers respected
// ------------------------------------------------------------
func TestGenerateOutfit_MaxUpperLayersLimit(t *testing.T) {

    items := []models.ClothingItem{
        {ID: 1, CategoryID: 1, Warmth: 1, Layer: intPtr(0)},
        {ID: 2, CategoryID: 1, Warmth: 1, Layer: intPtr(1)},
        {ID: 3, CategoryID: 1, Warmth: 1, Layer: intPtr(2)},
    }

    rule := &models.WeatherRule{
        RequiredWarmth: 10,
        MaxUpperLayers: 2,
    }

    outfit := GenerateOutfit(items, rule)

    if len(outfit.Upperwear) != 2 {
        t.Errorf("Expected 2 layers max, got %d", len(outfit.Upperwear))
    }
}
