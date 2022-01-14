package http

func (h *Http) Routes() {
	r := h.router.Group("/api")

	r.Post("/user", h.userHandler.CreateUser())
	r.Post("/weight", h.weightHandler.CreateWeight())
}
