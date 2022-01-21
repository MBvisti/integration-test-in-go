package http

import "github.com/gofiber/fiber/v2"

func (h *Http) Routes() {
	r := h.router.Group("/api")
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	r.Post("/user", h.userHandler.CreateUser())
	r.Post("/weight", h.weightHandler.CreateWeight())
}
