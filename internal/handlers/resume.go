package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func ResumeHandler(c *fiber.Ctx) error {
	return c.Render("internal/templates/resume", fiber.Map{
		"Title":     "Resume - Gsvd",
		"Canonical": "resume",
	}, "internal/templates/layouts/main")
}
