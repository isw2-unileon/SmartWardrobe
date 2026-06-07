# Database

## Database Engine

PostgreSQL (Supabase)

## Main Tables

### clothing_items

Stores user garments.

Fields:

- id
- user_id
- type_id
- color_id
- style_id
- image_url

### master_types

Stores garment types.

Examples:

- T-shirt
- Hoodie
- Jacket
- Jeans

### master_colors

Stores available colors.

### master_styles

Stores available styles.

Examples:

- Casual
- Formal
- Sporty

### master_categories

Stores garment categories.

Examples:

- Upperwear
- Bottomwear
- Footwear
- Outerwear

## Relationships

clothing_items

- belongs to master_types
- belongs to master_colors
- belongs to master_styles

master_types

- belongs to master_categories