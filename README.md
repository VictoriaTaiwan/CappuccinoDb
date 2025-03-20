☕ CappuchinoDB is a food database CRUD API, it contains ingredients (e.g., milk, sugar) and recipes for meals.

📌 API Endpoints

***

🥕 Ingredients.

| Method | Endpoint | Description |
| :---         |     :---:      |          ---: |
| POST   | /api/products/post      | Add a new ingredient to the database     |
| PUT     | /api/products/update       | Update an existing ingredient (requires id)      |
| GET     | /api/products/{name}       | Find an ingredient by name      |
| DELETE     | /api/products/{id}       | Delete an ingredient by id      |

***

🍲 Recipes.

| Method | Endpoint | Description |
| :---         |     :---:      |          ---: |
| POST   | /api/recipes/post     | Add a new recipe to the database     |
| PUT     | git diff       | Update an existing recipe (requires id)      |
| GET     | git diff       | Find recipes based on calorie count      |
| DELETE     | git diff       | Delete a recipe by id      |
