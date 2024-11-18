package main

import (
	"api/config"
	"api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{
		DisableStartupMessage:    true,
		Prefork:                  false,
		Concurrency:              1 * 1024 * 1024,
		DisableHeaderNormalizing: false,
	})

	config.Connect()
	defer config.Close()

	routes.Setup(app)

	app.Listen(":4003")
}
