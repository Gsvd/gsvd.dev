package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func ResumeHandler(c *fiber.Ctx) error {
	return c.Render("resume", fiber.Map{
		"Title": "Resume - Gsvd",
	}, "layouts/main")
}
