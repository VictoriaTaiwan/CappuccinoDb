package repositories

import (
	"context"
	"encoding/json"

	"cappuchinodb.com/main/app/data"
	"cappuchinodb.com/main/app/data/models"
)

func CreateRecipe(recipe models.Recipe) error {
	db, err := data.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close(context.Background())

	info := recipe.Information
	jsonIngredients, _ := json.Marshal(recipe.Ingredients)

	_, err = db.Exec(context.Background(),
	"INSERT INTO recipes (name, calories, unit_name, image_src, ingredients, instructions) VALUES ($1, $2, $3, $4, $5, $6)",
	info.Name, info.Calories, info.UnitName, info.ImageSrc, jsonIngredients, recipe.Instructions)
	if err != nil {
		return err
	}

	return nil
}

// Find all recipes by calories
func FilterRecipesByCalories(calories float64) ([]models.Recipe, error) {
	db, err := data.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(), 
	"SELECT id, name, calories, unit_name, image_src, ingredients, instructions FROM recipes WHERE calories=$1", 
	calories)

	var recipes []models.Recipe
	var info models.Product
	for rows.Next() {
		var recipe models.Recipe
		recipe.Information = info
		if err := rows.Scan(&info.ID, &info.Name, &info.Calories, &info.UnitName, &info.ImageSrc, &recipe.Ingredients, &recipe.Instructions); err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

// Update recipe's ingredients by recipe's ID
func UpdateRecipe(recipeID int, recipe models.Recipe) error {
	db, err := data.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close(context.Background())

	_, err = db.Exec(context.Background(),
		"UPDATE recipes SET name=$1, instructions=$2 WHERE id=$3",
		recipe.Information.Name, recipe.Instructions, recipeID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteRecipe(recipeID int) error {
	db, err := data.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close(context.Background())

	_, err = db.Exec(context.Background(), "DELETE FROM recipes WHERE id=$1", recipeID)
	if err != nil {
		return err
	}

	return nil
}
