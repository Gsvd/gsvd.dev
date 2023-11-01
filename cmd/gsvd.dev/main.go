package main

import (
	"log"
	"net/http"

	embeded "github.com/Gsvd/gsvd.dev"
	"github.com/Gsvd/gsvd.dev/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.NewFileSystem(http.FS(embeded.TemplateFiles), ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use("/css", filesystem.New(filesystem.Config{
		Root:       http.FS(embeded.DistFiles),
		PathPrefix: "dist/css",
		Browse:     true,
	}))
	app.Use("/js", filesystem.New(filesystem.Config{
		Root:       http.FS(embeded.PublicFiles),
		PathPrefix: "public/js",
		Browse:     true,
	}))
	app.Use("/images", filesystem.New(filesystem.Config{
		Root:       http.FS(embeded.PublicFiles),
		PathPrefix: "public/images",
		Browse:     true,
	}))
	app.Use("/fonts", filesystem.New(filesystem.Config{
		Root:       http.FS(embeded.PublicFiles),
		PathPrefix: "public/fonts",
		Browse:     true,
	}))

	app.Get("/", handlers.HomeHandler)
	app.Get("/blog", handlers.BlogHandler)
	app.Get("/blog/:title", handlers.BlogPostHandler)
	app.Get("/resume", handlers.ResumeHandler)

	log.Fatal(app.Listen(":3000"))
}
