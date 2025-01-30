package handlers

import (
	"github.com/Gsvd/gsvd.dev/internal"
	"github.com/gofiber/fiber/v2"
)

func HomeHandler(c *fiber.Ctx) error {
	articlesMetadata, err := internal.LoadArticlesMetadata()
	if err != nil {
		panic(err)
	}
	if len(articlesMetadata) > 5 {
		articlesMetadata = articlesMetadata[:5]
	}
	return c.Render("templates/index", fiber.Map{
		"Title":     "Gsvd - People-Focused Software Developer",
		"Articles":  articlesMetadata,
		"Canonical": "",
	}, "templates/layouts/main")
}
