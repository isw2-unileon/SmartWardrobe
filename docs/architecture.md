# Architecture

## Overview

Smart Wardrobe follows a client-server architecture.

The system is divided into three main layers:

1. Frontend (Next.js)
2. Backend API (Go + Gin)
3. Database and Authentication (Supabase/PostgreSQL)

The frontend communicates with the backend through REST APIs. The backend processes business logic, accesses the database, and integrates external services such as Open-Meteo, Remove.bg and CLIP. :contentReference[oaicite:1]{index=1}

## Frontend

The frontend is developed using:

- Next.js
- React
- TypeScript

Responsibilities:

- User authentication
- Clothing management
- Outfit generation requests
- Image uploads
- Displaying recommendations

Main folders:

- app/
- components/
- services/
- utils/

## Backend

The backend is implemented in Go using Gin.

Responsibilities:

- Business logic
- Database access
- Weather processing
- Outfit generation
- AI integration

Architecture pattern:

Controller → Service → Repository → Database

## Database

The database is PostgreSQL hosted in Supabase.

Main entities:

- Users
- Clothing Items
- Types
- Colors
- Styles
- Categories

## Storage

### Supabase Storage

Used to store garment images uploaded by users.

Uploaded images are stored in cloud storage and their URLs are saved in the database.

## External Services

### Open-Meteo

Used for weather forecasts.

### Remove.bg

Used for background removal from uploaded images.

### CLIP

Used for automatic clothing classification of:

- Garment type
- Garment color
- Garment style

## Communication Flow

1. The user interacts with the Next.js frontend.
2. The frontend sends HTTP requests to the Go backend.
3. The backend validates requests and executes business logic.
4. Data is retrieved or stored in PostgreSQL.
5. Images are stored in Supabase Storage.
6. External services may be invoked when necessary.
7. Results are returned to the frontend as JSON responses.