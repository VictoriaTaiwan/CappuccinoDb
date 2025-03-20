package main

import (
	"cappuchinodb.com/main/app/database"
	"cappuchinodb.com/main/app/pkg/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":8080")
}