package http

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type serverConfig interface {
	GetServerReadTimeOut() time.Duration
	GetServerWriteTimeOut() time.Duration
	GetServerPort() int64
}

type Http struct {
	router        *fiber.App
	serverPort    int64
	userHandler   userHandler
	weightHandler weightHandler
}

func NewHttp(
	cfg serverConfig, userHandler userHandler, weightHandler weightHandler) *Http {
	r := fiber.New(fiber.Config{
		ReadTimeout:  cfg.GetServerReadTimeOut(),
		WriteTimeout: cfg.GetServerWriteTimeOut(),
		AppName:      "Weight Tracking App",
	})
	return &Http{
		router:        r,
		serverPort:    cfg.GetServerPort(),
		userHandler:   userHandler,
		weightHandler: weightHandler,
	}
}

func (h *Http) SetupRoutes() {
	h.Routes()
}

func (h *Http) Start() error {
	h.SetupRoutes()
	return h.router.Listen(fmt.Sprintf(":%d", h.serverPort))
}

func (h *Http) GetRouter() *fiber.App {
	return h.router
}
