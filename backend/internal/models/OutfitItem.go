package models

type OutfitItem struct {
    Upperwear  []*ClothingItem
    Bottomwear *ClothingItem
    Footwear   *ClothingItem
}
