package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	embeded "github.com/Gsvd/gsvd.dev"
	"github.com/Gsvd/gsvd.dev/internal/models"
	"github.com/Gsvd/gsvd.dev/internal/services"
	"github.com/gernest/front"
	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/russross/blackfriday/v2"
)

func BlogHandler(c *fiber.Ctx) error {
	articlesMetadata, err := services.LoadMetadatas()
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
	var (
		filename = fmt.Sprintf("internal/content/%s.md", c.Params("title"))
		image    = "0.jpg"
	)

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

	if metadata.Image != "" {
		image = metadata.Image
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
		"Image":     image,
	}, "internal/templates/layouts/post")
}

func BlogLoadCommentsHandler(c *fiber.Ctx) error {
	articleId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid article Id")
	}

	comments, err := services.LoadComments(articleId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error loading comments")
	}

	return c.Render("internal/templates/partials/comments", fiber.Map{
		"Comments": comments,
	})
}

func BlogAddCommentHandler(c *fiber.Ctx) error {
	var (
		comment = &models.Comment{
			Username:  "Anonymous",
			CreatedAt: time.Now().UTC(),
		}
	)

	articleId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid article Id")
	}

	comment.PostId = articleId

	if value := c.FormValue("username"); value != "" && len(value) <= 16 {
		comment.Username = value
	}

	if value := c.FormValue("comment"); value != "" && len(value) <= 512 {
		comment.Comment = value
	} else if len(value) > 512 {
		comment.Comment = value[:509] + "..."
	}

	if err := services.SaveComment(comment); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error saving comment")
	}

	comment.CreatedAtFormatted = comment.CreatedAt.UTC().Format("January 2, 2006 at 3:04 PM")

	return c.Render("internal/templates/partials/comment", fiber.Map{
		"Comment":            comment.Comment,
		"Username":           comment.Username,
		"CreatedAtFormatted": comment.CreatedAtFormatted,
	})
}
