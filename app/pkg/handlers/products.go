package handlers

import (
	"strconv"

	"cappuchinodb.com/main/app/database"
	"cappuchinodb.com/main/app/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func SetupProductsRoutes(router fiber.Router) {
	products:= router.Group("/products")

	products.Get("/", GetProductsHandler)
	products.Post("/post", PostProductHandler)
	products.Put("/update", UpdateProductHandler)
	products.Delete("/delete", DeleteProductHandler)
}

func GetProductsHandler(c *fiber.Ctx) error {
	name := c.Query("name")

	product, err:= database.GetProduct(name)
	if err!=nil{
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No products found", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Product found", "data": product})
}

func PostProductHandler(c *fiber.Ctx) error {
	var product models.Product
	
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Error decoding JSON", "data": nil})
	}

	err = database.CreateProduct(product)
	if err!=nil{
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Product creation error", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Product created", "data": product})
}

func UpdateProductHandler(c *fiber.Ctx) error {
	var product models.Product
	
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Error decoding JSON", "data": nil})
	}

	err = database.UpdateProduct(product.ID, product)
	if err!=nil{
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Product update error", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Product updated", "data": product})
}

func DeleteProductHandler(c *fiber.Ctx) error {
	id := c.Query("id")
	
	idNum, convErr:=strconv.Atoi(id)
	if convErr!=nil{
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Id is invalid", "data": nil})
	}
	
	dbErr:= database.DeleteProduct(idNum)
	if dbErr!=nil{
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Product deletion error", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Product deleted", "data": id})
}
