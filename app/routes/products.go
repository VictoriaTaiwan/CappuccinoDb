package routes

import "github.com/gofiber/fiber/v2"

func SetupProductsRoutes(router fiber.Router) {
	products:= router.Group("/products")

	products.Get("/", GetProductsHandler)
	products.Post("/", PostProductHandler)
	products.Put("/", PutProductHandler)
	products.Delete("/", DeleteProductHandler)
}

func GetProductsHandler(c *fiber.Ctx) error {
	id := c.Query("id")
	calories := c.Query("calories")
	name := c.Query("name")

	info := "Product ID: " + id + ", calories: " + calories + ", name: " + name
	return c.SendString("Get products " + info)
}

func PostProductHandler(c *fiber.Ctx) error {
	return c.SendString("Post")
}

func PutProductHandler(c *fiber.Ctx) error {
	return c.SendString("Put")
}

func DeleteProductHandler(c *fiber.Ctx) error {
	return c.SendString("Delete")
}
