package router

import (
	"github.com/gofiber/fiber/v2"
	"placeholder-image/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/:resolution", handlers.GetImage)
}
