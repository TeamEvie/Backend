package main

import (
	"github.com/gominima/middlewares"
	"github.com/TeamEvie/Backend/routes"
	"github.com/gominima/cors"
	"github.com/gominima/minima"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := minima.New()
	crs := cors.New()
	app.UseRouter(routes.Router())
	app.UseRaw(middleware.Logger)
	app.Use(crs.NewCors(cors.Options{
		AllowedOrigins:   []string{"https://localhost:3000"},
		AllowCredentials: true,
		Debug:            false,
	}))
	app.Get("/", func(res *minima.Response, req *minima.Request) {
		res.OK().Send("Hello World")
	})
	app.Listen(":3000")
}
