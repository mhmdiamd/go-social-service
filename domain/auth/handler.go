package auth

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
    svc : svc,
  }
}

func (h handler) sendOtp(ctx *fiber.Ctx) error {
  req := SendOtpRequestPayload{}

  if err := ctx.BodyParser(&req); err != nil {
    return infrafiber.NewResponse(
      infrafiber.WithError(err),
      infrafiber.WithMessage(err.Error()),
      infrafiber.WithHttpCode(http.StatusBadRequest),
    ).Send(ctx)
  } 

  if err := h.svc.sendOtp(ctx.UserContext(), req); err != nil {
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

func (h handler) verifyOtp(ctx *fiber.Ctx) error {
  req := VerifyOtpRequestPayload{}

  if err := ctx.BodyParser(&req); err != nil {
    return infrafiber.NewResponse(
      infrafiber.WithError(err),
      infrafiber.WithMessage(err.Error()),
      infrafiber.WithHttpCode(http.StatusBadRequest),
    ).Send(ctx)
  } 

  otp_id, err := h.svc.verifyOtp(ctx.UserContext(), req); 
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
      "otp_id" : otp_id,
    }),
    infrafiber.WithMessage("success verify otp"),
    infrafiber.WithHttpCode(http.StatusOK),
  ).Send(ctx)
}

func (h handler) register(ctx *fiber.Ctx) error {
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
  if err := h.svc.register(ctx.UserContext(), req); err != nil {
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
  token, err := h.svc.login(ctx.UserContext(), req)
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
      "access_token" : token,
    }),
  ).Send(ctx)
}


