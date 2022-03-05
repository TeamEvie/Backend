package routes

import (
	"github.com/TeamEvie/Backend/routes/auth"
	"github.com/TeamEvie/Backend/routes/images"
	"github.com/gominima/minima"
)

func Router() *minima.Router {
	/* Create a new router */
	rt := minima.NewRouter()
	/* Define the routes */
	rt.Get("/test/:id", TestRoute())
	rt.Post("/sharex", images.UploadShareX())
	rt.Get("/sharex/sxcu", images.GenSXCU())
	rt.Get("/auth/github", auth.GitHubAuth())
	/* Return the router */
	return rt
}
