package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	embeded "github.com/Gsvd/gsvd.dev"
	"github.com/Gsvd/gsvd.dev/internal/handlers"
	"github.com/Gsvd/gsvd.dev/internal/stats"
	"github.com/Gsvd/gsvd.dev/internal/store"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	store.Init()
	stats.Init()

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
	app.Use("/stats", basicauth.New(basicauth.Config{
		Users: map[string]string{
			os.Getenv("STATS_USER"): os.Getenv("STATS_PASSWORD"),
		},
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
	app.Get("/blog/:id/comments", handlers.BlogLoadCommentsHandler)
	app.Post("/blog/:id/comment", handlers.BlogAddCommentHandler)
	app.Get("/resume", handlers.ResumeHandler)
	app.Get("/contact", handlers.ContactHandler)
	app.Get("/stats", func(c *fiber.Ctx) error {
		return c.SendFile(filepath.Join(os.Getenv("STATS_DIRECTORY"), os.Getenv("STATS_FILE")))
	})

	log.Fatal(app.Listen(":3000"))
}
