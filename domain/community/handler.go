package community

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	infrafiber "github.com/mhmdiamd/go-social-service/infra/fiber"
	"github.com/mhmdiamd/go-social-service/infra/response"
)

type handler struct {
	svc Service
}

func newHandler(svc Service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) GetAll(ctx *fiber.Ctx) error {
	var req ListCommunityRequestPayload

	if err := ctx.QueryParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithError(response.ErrorBadRequest),
			infrafiber.WithMessage("invalid payload"),
		).Send(ctx)
	}

	community, err := h.svc.GetAll(ctx.UserContext(), req)
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
		infrafiber.WithMessage("success delete by id"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(map[string]interface{}{
			"community": community,
		}),
		infrafiber.WithQuery(req),
	).Send(ctx)
}

func (h handler) GetById(ctx *fiber.Ctx) error {

	communityId, err := ctx.ParamsInt("id")

	if err != nil {

		fmt.Println(err)
		return infrafiber.NewResponse(
			infrafiber.WithMessage("id not valid"),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	community, err := h.svc.GetById(ctx.UserContext(), communityId)
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
		infrafiber.WithMessage("success delete by id"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(map[string]interface{}{
			"data": community,
		}),
	).Send(ctx)
}

func (h handler) Create(ctx *fiber.Ctx) error {

	parsedId := uuid.MustParse(ctx.Locals("PUBLIC_ID").(string))

	newCommunity := CreateCommunityRequestPayload{
		UserPublicId: parsedId,
	}

	if err := ctx.BodyParser(&newCommunity); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithError(err),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	err := h.svc.CreateCommunity(ctx.UserContext(), newCommunity)
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
		infrafiber.WithMessage("success create new community"),
		infrafiber.WithHttpCode(http.StatusOK),
	).Send(ctx)
}

func (h handler) UpdateById(ctx *fiber.Ctx) error {

	id, _ := ctx.ParamsInt("id")

	file, _ := ctx.FormFile("photo")
	name := ctx.FormValue("name")
	description := ctx.FormValue("description")
	externalCategories := ctx.FormValue("external_categories")

	categoryCommunityId := ctx.FormValue("category_community_id")
	intCategoryCommunityId, _ := strconv.Atoi(categoryCommunityId)

	req := UpdateCommunityRequestPayload{
		Id:                  id,
		Name:                name,
		Description:         description,
		ExternalCategories:  externalCategories,
		CategoryCommunityID: intCategoryCommunityId,
		Logo:                file,
	}

	err := h.svc.UpdateById(ctx.UserContext(), uuid.New(), req)
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
		infrafiber.WithMessage("success update by id"),
		infrafiber.WithHttpCode(http.StatusOK),
	).Send(ctx)
}

func (h handler) DeleteById(ctx *fiber.Ctx) error {

	communityId, err := ctx.ParamsInt("id")

	if err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("id not valid"),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	err = h.svc.DeleteById(ctx.UserContext(), communityId)
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
		infrafiber.WithMessage("success delete by id"),
		infrafiber.WithHttpCode(http.StatusOK),
	).Send(ctx)
}
