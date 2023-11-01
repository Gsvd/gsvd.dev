package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func ResumeHandler(c *fiber.Ctx) error {
	return c.Render("templates/resume", fiber.Map{
		"Title": "Resume - Gsvd",
	}, "templates/layouts/main")
}
