package handlers

import (
	"github.com/Gsvd/gsvd.dev/internal/services"
	"github.com/gofiber/fiber/v2"
)

func HomeHandler(c *fiber.Ctx) error {
	articlesMetadata, err := services.LoadMetadatas()
	if err != nil {
		panic(err)
	}
	if len(articlesMetadata) > 5 {
		articlesMetadata = articlesMetadata[:5]
	}
	return c.Render("internal/templates/index", fiber.Map{
		"Title":     "Gsvd - People-Focused Software Developer",
		"Articles":  articlesMetadata,
		"Canonical": "",
	}, "internal/templates/layouts/main")
}
