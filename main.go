package main

import (
	"github.com/TeamEvie/Backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()
	routes.Router(app)
	app.Use(logger.New())
	app.Use(NotFound)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Listen(":3000")
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/404.html")
}
