package http

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type userService interface {
	New(name, sex, weightGoal, email string, age, height,
		activityLevel int) error
}

type userHandler struct {
	service userService
}

func NewUserHandler(service userService) *userHandler {
	return &userHandler{
		service: service,
	}
}

type NewUserRequest struct {
	Name          string `json:"name"`
	Sex           string `json:"sex"`
	WeightGoal    string `json:"weight_goal"`
	Email         string `json:"email"`
	Age           int    `json:"age"`
	Height        int    `json:"height"`
	ActivityLevel int    `json:"activity_level"`
}

func (h *userHandler) CreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := NewUserRequest{}
		if err := c.BodyParser(&req); err != nil {
			log.Print(err)
			return c.Status(http.StatusBadRequest).
				SendString("could not parse request body")
		}

		err := h.service.New(
			req.Name, req.Sex, req.WeightGoal, req.Email, req.Age, req.Height,
			req.ActivityLevel,
		)
		if err != nil {
			return c.Status(http.StatusInternalServerError).
				SendString(err.Error())
		}

		return c.SendStatus(http.StatusOK)
	}
}
