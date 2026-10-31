package controllers

import (
	"cariinkos/app/dto"
	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	serviceContracts "cariinkos/app/services/contracts"

	"github.com/goravel/framework/contracts/http"
)

type KosFacilityController struct {
	service serviceContracts.KosFacilityService
}

func NewKosFacilityController(service serviceContracts.KosFacilityService) *KosFacilityController {
	return &KosFacilityController{service: service}
}

func (c *KosFacilityController) Index(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	facilities, err := c.service.GetAllByKosID(id)
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "Failed to retrieve kos facilities"})
	}
	return ctx.Response().Success().Json(http.Json{"data": facilities})
}

func (c *KosFacilityController) Store(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	userID, err := facades.Auth(ctx).ID()
	if err != nil || userID == "" {
		return ctx.Response().Status(http.StatusUnauthorized).Json(http.Json{"message": "Unauthorized"})
	}
	var request dto.AttachKosFacilityRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return ctx.Response().Status(http.DefaultAbortStatus).Json(http.Json{"message": "Invalid request body"})
	}
	validator, err := ctx.Request().Validate(map[string]any{"facility_id": "required|string"})
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "Validation failed"})
	}
	if validator.Fails() {
		return ctx.Response().Status(http.StatusUnprocessableEntity).Json(http.Json{"message": "Validation failed", "errors": validator.Errors().All()})
	}
	if err := c.service.Attach(userID, id, request.FacilityID); err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Status(http.StatusCreated).Json(http.Json{"message": "Facility attached to kos successfully"})
}

func (c *KosFacilityController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	userID, err := facades.Auth(ctx).ID()
	if err != nil || userID == "" {
		return ctx.Response().Status(http.StatusUnauthorized).Json(http.Json{"message": "Unauthorized"})
	}
	if err := c.service.Detach(userID, id, ctx.Request().Route("facility_id")); err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"message": "Facility detached from kos successfully"})
}
