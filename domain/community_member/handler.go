package communitymember

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

func newHandler(svc service) handler {
  return handler {
    svc : svc,
  }
}

func (h *handler) GetAllMemberByCommunityId(ctx *fiber.Ctx) error {

  var req CommunityMemberListRequestPayload

  community_id, err := ctx.ParamsInt("community_id", 0)
  fmt.Println(community_id)

  if err != nil {
    return infrafiber.NewResponse(
      infrafiber.WithMessage(err.Error()),
      infrafiber.WithError(response.ErrorBadRequest),
    ).Send(ctx)
  }

  if community_id == 0 {
    return infrafiber.NewResponse(
      infrafiber.WithMessage("invalid payload"),
      infrafiber.WithError(response.ErrorBadRequest),
    ).Send(ctx)
  }

  if err := ctx.QueryParser(&req); err != nil {
    return infrafiber.NewResponse(
      infrafiber.WithMessage("Invalid query payload"),
      infrafiber.WithError(err),
    ).Send(ctx)
  }

  members, err := h.svc.GetAllMemberByCommunityId(ctx.UserContext(), community_id, req); 
  if err != nil {
    myErr, ok := response.ErrorMapping[err.Error()];   
    if !ok {
      err = response.ErrorGeneral
    }

    return infrafiber.NewResponse(
      infrafiber.WithError(myErr),
      infrafiber.WithMessage(myErr.Error()),
    ).Send(ctx)
  }

  pagination := req.GenerateDefaultValue()
  return infrafiber.NewResponse(
    infrafiber.WithQuery(pagination),
    infrafiber.WithHttpCode(http.StatusOK),
    infrafiber.WithMessage("success get list member"),
    infrafiber.WithPayload(map[string]interface{} {
      "data" : members,
    }),
  ).Send(ctx)
}

func (h *handler) CreateNewMember(ctx *fiber.Ctx) error {
  var req AddCommunityMemberRequestPayload

  if err := ctx.BodyParser(&req); err != nil {
    return infrafiber.NewResponse(
      infrafiber.WithError(response.ErrorBadRequest),
      infrafiber.WithMessage("invalid payload"),
    ).Send(ctx)
  }

  err := h.svc.AddMember(ctx.UserContext(), req)
  if err != nil {
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
    infrafiber.WithHttpCode(http.StatusOK),
    infrafiber.WithMessage("success create new member"),
  ).Send(ctx)
}

func (h *handler) DeleteMemberById(ctx *fiber.Ctx) error {
  memberId := ctx.Params("member_id", "")
  communityId, err := ctx.ParamsInt("community_id", 0)

  if memberId == "" || err != nil || communityId == 0 {
    return infrafiber.NewResponse(
      infrafiber.WithError(response.ErrorBadRequest),
      infrafiber.WithMessage("invalid payload"),
    ).Send(ctx)
  }

  err = h.svc.DeleteCommunityMember(ctx.UserContext(), memberId, communityId);
  if err != nil {
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
    infrafiber.WithMessage("success delete member"),
    infrafiber.WithHttpCode(http.StatusOK),
  ).Send(ctx)
}

func (h *handler) UpdateMemberById(ctx *fiber.Ctx) error {
  var req UpdateCommunityMemberRequestPayload

  memberId := ctx.Params("member_id", "")
  communityId, err := ctx.ParamsInt("community_id", 0)

  if err != nil || communityId == 0 || memberId == ""{
    return infrafiber.NewResponse(
      infrafiber.WithError(response.ErrorBadRequest),
      infrafiber.WithMessage("invalid payload"),
    ).Send(ctx)
  }

  if err := ctx.BodyParser(&req); err != nil {
    return infrafiber.NewResponse(
      infrafiber.WithError(response.ErrorBadRequest),
      infrafiber.WithMessage("invalid payload"),
    ).Send(ctx)
  }

  req.UserPublicId = memberId
  req.CommunityId = communityId 
  err = h.svc.Update(ctx.UserContext(), req)
  fmt.Println(err)

  if err != nil {
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
    infrafiber.WithHttpCode(http.StatusOK),
    infrafiber.WithMessage("success update member"),
  ).Send(ctx)
}

func (h *handler) KickMemberById(ctx *fiber.Ctx) error {

	parsedId := uuid.MustParse(ctx.Locals("PUBLIC_ID").(string))

  memberId := ctx.Params("member_id", "")
  communityId, err := ctx.ParamsInt("community_id", 0)

  if err != nil || communityId == 0 || memberId == ""{
    return infrafiber.NewResponse(
      infrafiber.WithError(response.ErrorBadRequest),
      infrafiber.WithMessage("invalid payload"),
    ).Send(ctx)
  }

  err = h.svc.KickMember(ctx.UserContext(), parsedId.String(), memberId, communityId)
  if err != nil {
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
    infrafiber.WithHttpCode(http.StatusOK),
    infrafiber.WithMessage("success kick member"),
  ).Send(ctx)
}

