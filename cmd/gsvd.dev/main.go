package main

import (
	"log"

	"github.com/Gsvd/gsvd.dev/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/css", "./static/css")
	app.Static("/js", "./public/js")
	app.Static("/images", "./public/images")
	app.Static("/fonts", "./public/fonts")

	app.Get("/", handlers.HomeHandler)
	app.Get("/blog", handlers.BlogHandler)
	app.Get("/blog/:title", handlers.BlogPostHandler)
	app.Get("/resume", handlers.ResumeHandler)

	log.Fatal(app.Listen(":3000"))
}
