package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	infrafiber "github.com/mhmdiamd/go-social-service/infra/fiber"
	"github.com/mhmdiamd/go-social-service/infra/response"
)

type handler struct {
	Svc Service
}

func newHandler(svc Service) handler {
	return handler{
		Svc: svc,
	}
}

func (h handler) SendOtp(ctx *fiber.Ctx) error {
	req := SendOtpRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithError(err),
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	if err := h.Svc.SendOtp(ctx.UserContext(), req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]

		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithError(myErr),
			infrafiber.WithHttpCode(myErr.HttpCode),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithMessage("success send otp"),
		infrafiber.WithHttpCode(http.StatusOK),
	).Send(ctx)
}

func (h handler) VerifyOtp(ctx *fiber.Ctx) error {
	req := VerifyOtpRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithError(err),
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	otp_id, err := h.Svc.VerifyOtp(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]

		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithError(myErr),
			infrafiber.WithHttpCode(myErr.HttpCode),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithPayload(map[string]string{
			"otp_id": otp_id,
		}),
		infrafiber.WithMessage("success verify otp"),
		infrafiber.WithHttpCode(http.StatusOK),
	).Send(ctx)
}

func (h handler) Register(ctx *fiber.Ctx) error {
	//  Assign payload
	var req = RegisterRequestPayload{}

	//  Body parsing
	if err := ctx.BodyParser(&req); err != nil {
		myErr := response.ErrorBadRequest

		return infrafiber.NewResponse(
			infrafiber.WithError(myErr),
			infrafiber.WithMessage(myErr.Error()),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)

	}

	//  Hit service
	if err := h.Svc.Register(ctx.UserContext(), req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]

		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithError(myErr),
			infrafiber.WithMessage("register fail"),
		).Send(ctx)
	}

	// Return
	return infrafiber.NewResponse(
		infrafiber.WithMessage("register success"),
		infrafiber.WithHttpCode(http.StatusOK),
	).Send(ctx)
}

func (h handler) login(ctx *fiber.Ctx) error {
	var req = LoginRequestPayload{}

	// Parsing body
	if err := ctx.BodyParser(&req); err != nil {
		myErr := response.ErrorBadRequest

		return infrafiber.NewResponse(
			infrafiber.WithError(myErr),
			infrafiber.WithMessage("login fail"),
		).Send(ctx)
	}

	// Hit service
	token, err := h.Svc.login(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]

		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithError(myErr),
			infrafiber.WithMessage("login fail"),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithMessage("login success"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(map[string]string{
			"access_token": token,
		}),
	).Send(ctx)
}
