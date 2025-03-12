package handlers

import (
	"strconv"

	"cappuchinodb.com/main/app/database"
	"cappuchinodb.com/main/app/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func SetupRecipesRoutes(router fiber.Router) {
	recipes := router.Group("/recipes")

	recipes.Get("/", GetRecipesHandler)
	recipes.Post("/post", PostRecipeHandler)
	recipes.Put("/update", PutRecipeHandler)
	recipes.Delete("/delete", DeleteRecipeHandler)
}

func GetRecipesHandler(c *fiber.Ctx) error {
	calories := c.Query("calories")
	caloriesNum, convErr := strconv.ParseFloat(calories, 64)

	if convErr!=nil{
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Calories number is invalid", "data": nil})
	}

	recipes, dbErr:= database.FilterRecipesByCalories(caloriesNum)
	if dbErr!=nil{
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No recipes found", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Recipes found", "data": recipes})
}

func PostRecipeHandler(c *fiber.Ctx) error {
	var recipe models.Recipe
	
	err := c.BodyParser(&recipe)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Error decoding JSON", "data": nil})
	}

	err = database.CreateRecipe(recipe)
	if err!=nil{
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Recipe creation error", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Recipe created", "data": recipe})
}

func PutRecipeHandler(c *fiber.Ctx) error {
	var recipe models.Recipe
	
	err := c.BodyParser(&recipe)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Error decoding JSON", "data": nil})
	}

	err = database.UpdateRecipe(recipe.Information.ID, recipe)
	if err!=nil{
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Recipe update error", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Recipe updated", "data": recipe})
}

func DeleteRecipeHandler(c *fiber.Ctx) error {
	id := c.Query("id")
	
	idNum, convErr:=strconv.Atoi(id)
	if convErr!=nil{
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Id is invalid", "data": nil})
	}
	
	dbErr:= database.DeleteRecipe(idNum)
	if dbErr!=nil{
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Recipe deletion error", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Recipe deleted", "data": id})
}

