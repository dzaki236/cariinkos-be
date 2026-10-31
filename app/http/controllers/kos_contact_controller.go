package controllers

import (
	"cariinkos/app/dto"
	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	serviceContracts "cariinkos/app/services/contracts"

	"github.com/goravel/framework/contracts/http"
)

type KosContactController struct {
	service serviceContracts.KosContactService
}

func NewKosContactController(service serviceContracts.KosContactService) *KosContactController {
	return &KosContactController{service: service}
}

func (c *KosContactController) Index(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	contact, err := c.service.GetByKosID(id)
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "Failed to retrieve kos contacts"})
	}
	return ctx.Response().Success().Json(http.Json{"data": contact})
}

func (c *KosContactController) Store(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	userID, err := facades.Auth(ctx).ID()
	if err != nil || userID == "" {
		return ctx.Response().Status(http.StatusUnauthorized).Json(http.Json{"message": "Unauthorized"})
	}
	var request dto.CreateKosContactRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return ctx.Response().Status(http.DefaultAbortStatus).Json(http.Json{"message": "Invalid request body"})
	}
	validator, err := ctx.Request().Validate(map[string]any{"contact_name": "required|string|max:255", "phone": "required|string|max:50", "whatsapp": "nullable|string|max:50"})
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "Validation failed"})
	}
	if validator.Fails() {
		return ctx.Response().Status(http.StatusUnprocessableEntity).Json(http.Json{"message": "Validation failed", "errors": validator.Errors().All()})
	}
	contact, err := c.service.Create(userID, id, request)
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Status(http.StatusCreated).Json(http.Json{"data": contact})
}

func (c *KosContactController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	userID, err := facades.Auth(ctx).ID()
	if err != nil || userID == "" {
		return ctx.Response().Status(http.StatusUnauthorized).Json(http.Json{"message": "Unauthorized"})
	}
	var request dto.UpdateKosContactRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return ctx.Response().Status(http.DefaultAbortStatus).Json(http.Json{"message": "Invalid request body"})
	}
	validator, err := ctx.Request().Validate(map[string]any{"contact_name": "required|string|max:255", "phone": "required|string|max:50", "whatsapp": "nullable|string|max:50"})
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "Validation failed"})
	}
	if validator.Fails() {
		return ctx.Response().Status(http.StatusUnprocessableEntity).Json(http.Json{"message": "Validation failed", "errors": validator.Errors().All()})
	}
	contact, err := c.service.Update(userID, id, request)
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"data": contact})
}

func (c *KosContactController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	userID, err := facades.Auth(ctx).ID()
	if err != nil || userID == "" {
		return ctx.Response().Status(http.StatusUnauthorized).Json(http.Json{"message": "Unauthorized"})
	}
	if err := c.service.Delete(userID, id); err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"message": "Kos contact deleted successfully"})
}
