// main.go
package main

import (
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	})
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")
}
