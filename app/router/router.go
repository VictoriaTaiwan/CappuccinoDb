package router

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "cappuchinodb.com/main/app/routes"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api", logger.New())

    routes.SetupProductsRoutes(api)
	routes.SetupRecipesRoutes(api)
}

