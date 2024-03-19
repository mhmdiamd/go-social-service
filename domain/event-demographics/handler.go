package eventdemographics

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
  return handler {
    svc: svc,
  }
}


func (h handler) Create(ctx *fiber.Ctx) error {

  // name := ctx.FormValue("gender")
  // gender := ctx.FormValue("gender")
  // graduation := ctx.FormValue("graduation")
  // startAge := ctx.FormValue("start_age")
  // endAge := ctx.FormValue("end_age")
  //
  // newStartAge, err := strconv.Atoi(startAge);
  // if err != nil {
  //   err = errors.New("start age is invalid")
  // }
  //
  // newEndAge, err := strconv.Atoi(endAge);
  // if err != nil {
  //   err = errors.New("end age is invalid")
  // }
  //
  //
  // req := CreateEventDemographicsRequestPayload{
  //   Name: name,
  //   Gender: gender,
  //   StartAge:  newStartAge,
  //   EndAge: newEndAge,
  // }

  var req CreateEventDemographicsRequestPayload

  if err := ctx.BodyParser(&req); err != nil {
    return infrafiber.NewResponse(
      infrafiber.WithError(err),
      infrafiber.WithHttpCode(http.StatusBadRequest),
    ).Send(ctx)
  }

  err := h.svc.AddEventDemographics(ctx.UserContext(), req)
  if err != nil {
    myErr, ok := response.ErrorMapping[err.Error()];
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
    infrafiber.WithMessage("success create event demographics"),
    infrafiber.WithHttpCode(http.StatusOK),
  ).Send(ctx)
}

func (h handler) UpdateById(ctx *fiber.Ctx) error {
  var req UpdateEventDemographicsRequestPayload

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
    myErr, ok := response.ErrorMapping[err.Error()];
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
    infrafiber.WithMessage("success update event demographics"),
    infrafiber.WithHttpCode(http.StatusOK),
  ).Send(ctx)
}

func (h handler) GetAll(ctx *fiber.Ctx) error {

  res, err := h.svc.GetListEventDemographics(ctx.UserContext())
  if err != nil {
    myErr, ok := response.ErrorMapping[err.Error()];
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
    infrafiber.WithMessage("success get all event demographics"),
    infrafiber.WithHttpCode(http.StatusOK),
    infrafiber.WithPayload(map[string]interface{}{
      "event_demographies" : res,
    }),
  ).Send(ctx)
}

func (h handler) GetById(ctx *fiber.Ctx) error {
  id, _ := ctx.ParamsInt("id")

  res, err := h.svc.GetById(ctx.UserContext(), id);
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
    infrafiber.WithMessage("success get event demographics by id"),
    infrafiber.WithHttpCode(http.StatusOK),
    infrafiber.WithPayload(map[string]interface{}{
      "event_demographics" : res,
    }),
  ).Send(ctx)
}

func (h handler) DeleteById(ctx *fiber.Ctx) error {
  id, _ := ctx.ParamsInt("id")

  err := h.svc.DeleteById(ctx.UserContext(), id);
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
    infrafiber.WithMessage("success delete event demographics"),
    infrafiber.WithHttpCode(http.StatusOK),
  ).Send(ctx)
}
