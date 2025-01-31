package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	embeded "github.com/Gsvd/gsvd.dev"
	"github.com/Gsvd/gsvd.dev/internal/models"
	"github.com/Gsvd/gsvd.dev/internal/services"
	"github.com/gernest/front"
	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/russross/blackfriday/v2"
)

func BlogHandler(c *fiber.Ctx) error {
	articlesMetadata, err := services.LoadArticles()
	if err != nil {
		panic(err)
	}
	return c.Render("internal/templates/blog", fiber.Map{
		"Title":     "Blog Articles - Gsvd",
		"Articles":  articlesMetadata,
		"Canonical": "blog",
	}, "internal/templates/layouts/main")
}

func BlogPostHandler(c *fiber.Ctx) error {
	filename := fmt.Sprintf("internal/content/%s.md", c.Params("title"))

	fileContent, err := embeded.ContentFiles.ReadFile(filename)
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	m := front.NewMatter()
	m.Handle("---", front.YAMLHandler)
	f, body, err := m.Parse(bytes.NewReader(fileContent))
	if err != nil {
		panic(err)
	}

	metadata := &models.Metadata{}
	if err := mapstructure.Decode(f, metadata); err != nil {
		panic(err)
	}

	htmlContent := blackfriday.Run([]byte(body))

	article := models.Article{
		Metadata: *metadata,
		Content:  template.HTML(htmlContent),
	}

	return c.Render("internal/templates/post", fiber.Map{
		"Title":     metadata.Title + " - Gsvd",
		"Article":   article,
		"Canonical": "blog/" + metadata.Slug,
	}, "internal/templates/layouts/post")
}
