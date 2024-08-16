package event

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	infrafiber "github.com/mhmdiamd/go-social-service/infra/fiber"
	"github.com/mhmdiamd/go-social-service/infra/response"
)

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
	parsedId := uuid.MustParse(ctx.Locals("PUBLIC_ID").(string))

	event := CreateEventRequestPayload{
		UserPublicId: parsedId,
	}

	if err := ctx.BodyParser(&event); err != nil {
		fmt.Printf("%v : %v", err, "error from parser")
		return infrafiber.NewResponse(
			infrafiber.WithError(err),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	err := h.svc.Create(ctx.UserContext(), event)
	if err != nil {
		fmt.Println(err)
		myErr, ok := response.ErrorMapping[err.Error()]

		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithError(myErr),
			infrafiber.WithMessage(myErr.Error()),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithMessage("success create new event"),
		infrafiber.WithHttpCode(http.StatusOK),
	).Send(ctx)
}

func (h *handler) DeleteById(ctx *fiber.Ctx) error {
	return nil
}

func (h *handler) UpdateById(ctx *fiber.Ctx) error {
	return nil
}
