package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tsmurer/pizzashop-be/database"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the pizzashop")
	})

	app.Listen(":3000")
}
