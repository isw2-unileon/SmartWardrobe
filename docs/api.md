# API Documentation

## Base URL

```text
/api
```

---

## Colors

### Get all colors

```http
GET /getAllColors
```

Returns all available clothing colors.

### Response

```json
[
  {
    "id": 1,
    "name": "Black"
  }
]
```

---

## Types

### Get all garment types

```http
GET /getAllTypes
```

Returns all available garment types.

### Response

```json
[
  {
    "id": 1,
    "name": "T-shirt"
  }
]
```

---

## Styles

### Get all clothing styles

```http
GET /getAllStyles
```

Returns all available styles.

### Response

```json
[
  {
    "id": 1,
    "name": "Casual"
  }
]
```

---

## Clothing Items

### Get all user clothing items

```http
GET /clothingItems
```

Returns every garment belonging to the authenticated user.

---

### Get clothing item by ID

```http
GET /clothingItem/{id}
```

Returns a specific clothing item.

---

### Filter clothing items

```http
GET /clothingItem/filters
```

Query parameters:

| Parameter | Description |
|-----------|-------------|
| typeId | Filter by type |
| colorId | Filter by color |
| styleId | Filter by style |

Example:

```http
GET /clothingItem/filters?typeId=1&colorId=2
```

---

### Create clothing item

```http
POST /clothingItem
```

Request:

```json
{
  "type": {
    "id": 1
  },
  "color": {
    "id": 2
  },
  "style": {
    "id": 3
  },
  "imageUrl": "https://..."
}
```

---

### Update clothing item

```http
PUT /clothingItem/{id}
```

Updates an existing clothing item.

---

### Delete clothing item

```http
DELETE /clothingItem/{id}
```

Deletes a clothing item.

---

## Outfit Generation

### Generate outfit recommendation

```http
POST /generateOutfit/days
```

Request:

```json
{
  "city": "Madrid",
  "country": "Spain",
  "start_date": "2026-06-01",
  "end_date": "2026-06-07"
}
```

Response:

```json
{
  "weather": {
    "daily": {
      "time": ["2026-06-01"],
      "temperature_2m_max": [27],
      "temperature_2m_min": [16]
    }
  },
  "outfit": {
    "upperwear": {},
    "bottomwear": {},
    "footwear": {},
    "outerwear": {}
  }
}
```

---

## Background Removal

### Remove image background

```http
POST /removeBackground
```

Content-Type:

```text
multipart/form-data
```

Parameters:

| Field | Type |
|--------|------|
| file | image |

Returns a PNG image with the background removed.

---

## AI Clothing Classification

### Analyze garment image

```http
POST /clothing/analyze
```

Content-Type:

```text
multipart/form-data
```

Parameters:

| Field | Type |
|--------|------|
| image | image |

Response:

```json
{
  "color": "Black",
  "style": "Casual",
  "type": "Tshirt"
}
```

Description:

This endpoint uses the CLIP model to automatically classify an uploaded garment image.

The service analyzes the image and predicts:

- Garment type
- Garment color
- Garment style

These predictions are displayed to the user during the garment verification process and can be manually adjusted before saving the item.