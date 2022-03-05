package routes

import (
	"github.com/TeamEvie/Backend/routes/auth"
	"github.com/TeamEvie/Backend/routes/images"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	/* Define v1 Group */
	v1 := app.Group("/v1")
	/* Define the routes */
	v1.Post("/sharex", images.UploadShareX)
	v1.Get("/sharex/sxcu", images.GenSXCU)
	v1.Get("/auth/github", auth.GitHubAuth)
}
