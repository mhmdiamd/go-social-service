package eventcommite

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	infrafiber "github.com/mhmdiamd/go-social-service/infra/fiber"
	"github.com/mhmdiamd/go-social-service/infra/response"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h *handler) Create(ctx *fiber.Ctx) error {
	var req CreateEventCommitteeRequestPayload

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithError(err),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	err := h.svc.AddEventCommittee(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithError(myErr),
			infrafiber.WithMessage(myErr.Message),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithMessage("success create event committee"),
		infrafiber.WithHttpCode(http.StatusOK),
	).Send(ctx)
}

func (h *handler) UpdateById(ctx *fiber.Ctx) error {
	var req UpdateEventCommitteeRequestPayload

	id, _ := ctx.ParamsInt("id")
	req.Id = id

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithError(err),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	err := h.svc.UpdateById(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithError(myErr),
			infrafiber.WithMessage(myErr.Message),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithMessage("success update event committee"),
		infrafiber.WithHttpCode(http.StatusOK),
	).Send(ctx)
}

func (h *handler) GetAll(ctx *fiber.Ctx) error {
	res, err := h.svc.GetListEventCommittee(ctx.UserContext())
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithError(myErr),
			infrafiber.WithMessage(myErr.Message),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithMessage("success get all event committee"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(map[string]interface{}{
			"event_committee": res,
		}),
	).Send(ctx)
}

func (h *handler) GetById(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	res, err := h.svc.GetById(ctx.UserContext(), id)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]

		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithError(myErr),
			infrafiber.WithMessage(myErr.Message),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithMessage("success get event committe by id"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(map[string]interface{}{
			"event_committee": res,
		}),
	).Send(ctx)
}

func (h *handler) DeleteById(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	err := h.svc.DeleteById(ctx.UserContext(), id)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]

		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithError(myErr),
			infrafiber.WithMessage(myErr.Message),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithMessage("success delete event committe"),
		infrafiber.WithHttpCode(http.StatusOK),
	).Send(ctx)
}
