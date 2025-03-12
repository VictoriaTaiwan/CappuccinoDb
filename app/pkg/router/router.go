package router

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "cappuchinodb.com/main/app/pkg/handlers"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api", logger.New())

    handlers.SetupProductsRoutes(api)
	handlers.SetupRecipesRoutes(api)
}

