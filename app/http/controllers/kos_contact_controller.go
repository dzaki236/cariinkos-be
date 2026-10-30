package controllers

import (
	"cariinkos/app/dto"
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
	contact, err := c.service.GetByKosID(ctx.Request().Route("id"))
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "Failed to retrieve kos contacts"})
	}
	return ctx.Response().Success().Json(http.Json{"data": contact})
}

func (c *KosContactController) Store(ctx http.Context) http.Response {
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
	contact, err := c.service.Create(ctx.Request().Route("id"), request)
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Status(http.StatusCreated).Json(http.Json{"data": contact})
}

func (c *KosContactController) Update(ctx http.Context) http.Response {
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
	contact, err := c.service.Update(ctx.Request().Route("id"), request)
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"data": contact})
}

func (c *KosContactController) Destroy(ctx http.Context) http.Response {
	if err := c.service.Delete(ctx.Request().Route("id")); err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"message": "Kos contact deleted successfully"})
}
