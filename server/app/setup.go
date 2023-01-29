package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"placeholder-image/config"
	"placeholder-image/router"
)

func main() {
	port, err := config.LoadPORT()

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(recover.New())

	router.SetupRoutes(app)

	app.Listen(":" + port)

}
