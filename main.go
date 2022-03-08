package main

import (
	"github.com/TeamEvie/Backend/routes"
	"github.com/gominima/middlewares"
	"github.com/gominima/minima"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := minima.New()
	app.UseGroup(routes.Router())
	app.UseRaw(middleware.Logger)
	app.NotFound(NotFound())
	app.Get("/", func(res *minima.Response, req *minima.Request) {
		res.OK().Send("Hello, World!")
	})
	app.Listen(":3000")
}

func NotFound() minima.Handler {
	return func(res *minima.Response, req *minima.Request) { 
		res.Status(404).Render("./static/404.html", "") 
	}
}
