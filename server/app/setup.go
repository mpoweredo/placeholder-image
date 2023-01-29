package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.SetupRoutes(app)

	app.Listen(":" + port)

}
