package categorycomunity

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

func(h handler) CreateCategoryCommunity(ctx *fiber.Ctx) error {

  var req = CreateCategoryCommunityRequestPayload{}

  if err := ctx.BodyParser(&req); err != nil {
    return infrafiber.NewResponse(
      infrafiber.WithError(err),
      infrafiber.WithHttpCode(http.StatusBadRequest),
    ).Send(ctx)
  }

  err := h.svc.AddCategoryCommunity(ctx.UserContext(), req)
  if err != nil {
    myErr, ok := response.ErrorMapping[err.Error()]

    if !ok {
      myErr = response.ErrorGeneral
    }

    return infrafiber.NewResponse(
      infrafiber.WithError(myErr),
      infrafiber.WithMessage("fail create category community"),
    ).Send(ctx)
  }


  return infrafiber.NewResponse(
    infrafiber.WithMessage("success create category community"),
    infrafiber.WithHttpCode(http.StatusOK),
  ).Send(ctx)
}


