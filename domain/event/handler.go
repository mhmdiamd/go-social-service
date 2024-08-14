package event

import "github.com/gofiber/fiber/v2"

type handler struct {
	svc service
}

func newHandler(svc service) *handler {
	return &handler{
		svc: svc,
	}
}

func (h *handler) GetAll(ctx *fiber.Ctx) error {
	return nil
}

func (h *handler) GetById(ctx *fiber.Ctx) error {
	return nil
}

func (h *handler) Create(ctx *fiber.Ctx) error {
	return nil
}

func (h *handler) DeleteById(ctx *fiber.Ctx) error {
	return nil
}

func (h *handler) UpdateById(ctx *fiber.Ctx) error {
	return nil
}
