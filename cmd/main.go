package main

import "github.com/gofiber/fiber/v3"

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, GET!")
	})

	app.Post("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, POST!")
	})

	app.Put("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, PUT!")
	})

	app.Delete("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, DELETE!")
	})

	app.Listen(":3000")
}
