package routes

import "github.com/gofiber/fiber/v2"

func SetupRecipesRoutes(router fiber.Router) {
	recipes := router.Group("/recipes")

	recipes.Get("/", GetRecipesHandler)
	recipes.Post("/", PostRecipeHandler)
	recipes.Put("/", PutRecipeHandler)
	recipes.Delete("/", DeleteRecipeHandler)
}

func GetRecipesHandler(c *fiber.Ctx) error {
	//id := c.Params("id")
	return c.SendString("Get")
}

func PostRecipeHandler(c *fiber.Ctx) error {
	return c.SendString("Post")
}

func PutRecipeHandler(c *fiber.Ctx) error {
	return c.SendString("Put")
}

func DeleteRecipeHandler(c *fiber.Ctx) error {
	return c.SendString("Delete")
}