package main

import (
	"./routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")
	engine.Reload(true)
	app := fiber.New(&fiber.Settings{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/", routes.Index)
	app.Post("/v1/new", routes.Write)

	app.Listen(3000)
}
