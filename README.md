# Smart Wardrobe
The project will consist of a virtual closet that helps the user automatically generate outfits using the clothing items uploaded to the app.

The application itself allows the user to upload photos of their clothes, view the uploaded items, add specific attributes to them, and search using filters.

Furthermore, the user will be able to ask the app to generate an outfit for a specific day or a range of days, indicating the activity they will be doing or the type of clothes they feel like wearing. Using this information, the application will retrieve the forecast through a weather API in order to suggest the most appropriate outfit.


## Technologies
* **Backend:** Go + Gin API
* **Frontend:** Next.js
* **Deploy:** Render.com


## Project Structure
The repository will be divided into two main parts:

```text
/
├── backend/    # Go server source code
├── frontend/   # Next.js client source code
└── README.md