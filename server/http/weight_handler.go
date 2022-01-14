package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type weightService interface {
	New(weight, userId, bmr, dailyCaloricIntake int) error
}

type weightHandler struct {
	service weightService
}

func NewWeightHandler(service weightService) *weightHandler {
	return &weightHandler{
		service: service,
	}
}

type NewWeightRequest struct {
	Weight             int `json:"weight"`
	UserId             int `json:"user_id"`
	BMR                int `json:"bmr"`
	DailyCaloricIntake int `json:"daily_caloric_intake"`
}

func (h *weightHandler) CreateWeight() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req NewWeightRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).
				SendString("could not parse request body")
		}

		err := h.service.New(
			req.Weight, req.UserId, req.BMR, req.DailyCaloricIntake,
		)
		if err != nil {
			return c.Status(http.StatusInternalServerError).
				SendString(err.Error())
		}

		return c.SendStatus(http.StatusOK)
	}
}
