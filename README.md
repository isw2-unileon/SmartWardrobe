# Smart Wardrobe
Smart Wardrobe is a web application that helps users manage their personal wardrobe and automatically generate outfit recommendations based on weather conditions.

Users can upload images of their clothing items, classify them by type, color, and style, and store them in a virtual wardrobe. The application allows browsing, searching, modifying, and deleting clothing items through an intuitive interface.

The outfit generation feature creates personalized outfit suggestions using the user's available clothes and real-time weather forecasts. Users can generate outfits for the current day, the next day, or an entire week. The system analyzes the expected temperature and selects suitable garments from the wardrobe to create complete outfit combinations.

The project follows a client-server architecture with a React/Next.js frontend, a Go backend exposing REST endpoints, Supabase authentication and storage, and PostgreSQL database persistence.


## Technologies
* **Backend:** Go + Gin API
* **Frontend:** Next.js + React + Typescript + Css
* **Deploy:** Render.com
* **Authentication and Storage** Supabase
* **External APIs** Open-Meteo API (weather forecast) + Open-Meteo Geocoding API (city coordinates) + Remove.bg API

## Project Structure
The repository will be divided into two main parts:

```text
/
├── backend/    # Go server source code
├── frontend/   # Next.js client source code
└── README.md

```

## Installation and local setup
### 1. Clone the Repository

```text
git clone  <repo-url>/SmartWardrobe.git
cd SmartWardrobe

```

### 2. Configure Environment Variables 
#### Frontend

Create a .env.local file inside the frontend directory:

```text 
NEXT_PUBLIC_API_URL=http://localhost:8080

NEXT_PUBLIC_SUPABASE_URL=<your-supabase-url>
NEXT_PUBLIC_SUPABASE_PUBLISHABLE_KEY=<your-supabase-pulishable-key>

```

#### Backend

Create a .env file inside the backend directory:
```text
DATABASE_URL=<postgresql-connection-string>

SUPABASE_URL=<your-supabase-url>
SUPABASE_JWT_SECRET=<your-supabase-jwt-secret>

REMOVEBG_API_KEY=<your-removebg-api-key>

```

### 3. Install Dependencies
#### Frontend

```text
cd frontend
npm install

```

### 4. Run the Backend

From the backend directory:
```text
go run cmd/server/main.go
```
The backend will start on:
```text
http://localhost:8080
```

### 5. Run the Frontend

From the frontend directory:
```text
npm run dev
```
The frontend will start on:
```text
http://localhost:3000
```

### 6. Access the Application

Open:
```text
http://localhost:3000
```
and log in with a valid account to access the virtual wardrobe features.

## Execution of backend tests
Execution of all test
```text
go test ./...
```

Execution of all test with coverture
```text
go test -cover ./...
```

## Contributing
### 1. Branching Strategy
Each new feature, bug fix, or improvement should be developed in a separate branch created from main.

Branch naming convention:
```text
feature/<feature-number>-<feature-name>
bugfix/<bugfix-number>-<bugfix-name>
hotfix/<hotfix-number>-<gotfix-name>
```

### 2. Commit Convention
Commits should be short, descriptive and written in English.

### 3. Pull Requests
Before opening a Pull Request:
    1. Ensure the project builds successfully.
    2. Run all available tests.
    3. Verify that the new functionality works as expected.
    4. Resolve merge conflicts if any exist.

After review and approval, the Pull Request can be merged into the main branch.

## Documentation