package main

import (
	embeded "github.com/Gsvd/gsvd.dev"
	"github.com/Gsvd/gsvd.dev/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	engine := html.NewFileSystem(http.FS(embeded.TemplateFiles), ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Use("/css", filesystem.New(filesystem.Config{
		Root:       http.FS(embeded.DistFiles),
		PathPrefix: "dist/css",
		Browse:     false,
	}))
	app.Use("/images", filesystem.New(filesystem.Config{
		Root:       http.FS(embeded.PublicFiles),
		PathPrefix: "public/images",
		Browse:     false,
	}))
	app.Use("/fonts", filesystem.New(filesystem.Config{
		Root:       http.FS(embeded.PublicFiles),
		PathPrefix: "public/fonts",
		Browse:     false,
	}))
	app.Get("/sitemap.xml", func(c *fiber.Ctx) error {
		file, err := embeded.SiteMapFile.ReadFile("sitemap.xml")
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		c.Set("Content-Type", "application/xml")
		return c.Send(file)
	})

	app.Get("/", handlers.HomeHandler)
	app.Get("/blog", handlers.BlogHandler)
	app.Get("/blog/:title", handlers.BlogPostHandler)
	app.Post("/blog/:title/:id", handlers.BlogPostCommentHandler)
	app.Get("/resume", handlers.ResumeHandler)
	app.Get("/contact", handlers.ContactHandler)

	log.Fatal(app.Listen(":3000"))
}
