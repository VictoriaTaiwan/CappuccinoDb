☕ CappuchinoDB is a food database CRUD API, it contains ingredients (e.g., milk, sugar) and recipes for meals.

📌 API Endpoints

***

🥕 Ingredients.

| Method | Endpoint | Description |
| :---         |     :---:      |          ---: |
| POST   | /api/products/post      | Add a new ingredient to the database     |
| PUT     | /api/products/update       | Update an existing ingredient (requires id)      |
| GET     | /api/products/`name={name}`       | Find an ingredient by name      |
| DELETE     | /api/products/delete?`id={id}`       | Delete an ingredient by id      |

Example for POST request

```
	data := bytes.NewBufferString(`{"name":"sugar","calories":4,"unit_name":"grams","image_src":"https://example/sugar_image.png"}`)
	req, err := http.NewRequest("PUT", "http://localhost:8080/api/products/post", data)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
```

***

🍲 Recipes.

| Method | Endpoint | Description |
| :---         |     :---:      |          ---: |
| POST   | /api/recipes/post     | Add a new recipe to the database     |
| PUT     | /api/recipes/update       | Update an existing recipe (requires id)      |
| GET     | /api/recipes/`calories={calories} `      | Find recipes based on calorie count      |
| DELETE     | /api/recipes/delete?`id={id}`       | Delete a recipe by id      |


Example for DELETE request

```
	url := "http://localhost:8080/api/products/delete?id=1"

	req, err := http.NewRequest("DELETE", url, nil) 

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
```
