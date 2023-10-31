package handlers

import (
	"bytes"
	"fmt"
	"github.com/Gsvd/gsvd.dev/internal"
	"github.com/gernest/front"
	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"net/http"
	"os"
)

func BlogHandler(c *fiber.Ctx) error {
	articlesMetadata, err := internal.LoadArticlesMetadata()
	if err != nil {
		panic(err)
	}
	return c.Render("blog", fiber.Map{
		"Title":    "Blog Articles - Gsvd",
		"Articles": articlesMetadata,
	}, "layouts/main")
}

func BlogPostHandler(c *fiber.Ctx) error {
	filename := fmt.Sprintf("./articles/%s.md", c.Params("title"))

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	m := front.NewMatter()
	m.Handle("---", front.YAMLHandler)
	f, body, err := m.Parse(bytes.NewReader(fileContent))
	if err != nil {
		panic(err)
	}

	metadata := &internal.ArticleMetadata{}
	if err := mapstructure.Decode(f, metadata); err != nil {
		panic(err)
	}

	htmlContent := blackfriday.Run([]byte(body))

	article := internal.Article{
		Metadata: *metadata,
		Content:  template.HTML(htmlContent),
	}

	return c.Render("post", fiber.Map{
		"Title":   metadata.Title + " - Gsvd",
		"Article": article,
	}, "layouts/main")
}