# Deployment

## Overview

Smart Wardrobe is deployed using a cloud-based architecture composed of:

- Render (Frontend)
- Render (Backend)
- Supabase (Database and Authentication)

---

## Frontend Deployment

### Platform

Render

### Technology

- Next.js
- React
- TypeScript

### Build Command

```bash
npm install
npm run build
```

### Start Command

```bash
npm start
```

### Required Environment Variables

```env
NEXT_PUBLIC_SUPABASE_URL=
NEXT_PUBLIC_SUPABASE_ANON_KEY=
NEXT_PUBLIC_API_URL=
```

---

## Backend Deployment

### Platform

Render

### Technology

- Go
- Gin
- GORM

### Build Command

```bash
go build -o server ./cmd/server
```

### Start Command

```bash
./server
```

### Required Environment Variables

```env
DATABASE_URL=
NEXT_URL=
REMOVEBG_API_KEY=
ORT_LIB_PATH=
```

---

## Database Deployment

### Platform

Supabase

### Engine

PostgreSQL

### Responsibilities

- User storage
- Clothing item storage
- Master data storage
- Authentication support

---

## External Services

### Remove.bg

Used to remove image backgrounds before storing garments.

### Open-Meteo

Used to retrieve weather forecasts for outfit generation.

### CLIP

Used to classify:

- Garment type
- Garment color
- Garment style

from uploaded images.

---

## Deployment Workflow

1. Push code to GitHub.
2. Render detects repository changes.
3. Frontend and backend are rebuilt automatically.
4. Applications connect to Supabase.
5. Services become available through public URLs.