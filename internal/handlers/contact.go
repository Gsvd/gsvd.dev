package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func ContactHandler(c *fiber.Ctx) error {
	return c.Render("templates/contact", fiber.Map{
		"Title":     "Contact - Gsvd",
		"Canonical": "contact",
	}, "templates/layouts/main")
}
