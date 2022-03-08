package routes

import (
	"github.com/TeamEvie/Backend/routes/auth"
	"github.com/TeamEvie/Backend/routes/images"
	"github.com/gominima/minima"
)

func Router() *minima.Group {
	/* Define v1 Group */
	v1 := minima.NewGroup("/v1")
	/* Define the routes */
	v1.Post("/sharex", images.UploadShareX())
	v1.Get("/sharex/sxcu", images.GenSXCU())
	v1.Get("/auth/github", auth.GitHubAuth())

	return v1
}
