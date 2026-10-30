package controllers

import (
	"cariinkos/app/dto"
	"cariinkos/app/helpers"
	serviceContracts "cariinkos/app/services/contracts"

	"github.com/goravel/framework/contracts/http"
)

type FacilityController struct {
	service serviceContracts.FacilityService
}

func NewFacilityController(service serviceContracts.FacilityService) *FacilityController {
	return &FacilityController{service: service}
}

func (c *FacilityController) Index(ctx http.Context) http.Response {
	facilities, err := c.service.GetAllData()
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "Failed to retrieve facilities"})
	}
	return ctx.Response().Success().Json(http.Json{"data": facilities})
}

func (c *FacilityController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	facility, err := c.service.GetByID(id)
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"data": facility})
}

func (c *FacilityController) Store(ctx http.Context) http.Response {
	var request dto.CreateFacilityRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return ctx.Response().Status(http.DefaultAbortStatus).Json(http.Json{"message": "Invalid request body"})
	}
	validator, err := ctx.Request().Validate(map[string]any{
		"name": "required|string|max:255",
		"icon": "nullable|string|max:255",
	})
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "Validation failed"})
	}
	if validator.Fails() {
		return ctx.Response().Status(http.StatusUnprocessableEntity).Json(http.Json{"message": "Validation failed", "errors": validator.Errors().All()})
	}

	facility, err := c.service.Create(request)
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Status(http.StatusCreated).Json(http.Json{"data": facility})
}

func (c *FacilityController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var request dto.UpdateFacilityRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return ctx.Response().Status(http.DefaultAbortStatus).Json(http.Json{"message": "Invalid request body"})
	}
	validator, err := ctx.Request().Validate(map[string]any{
		"name": "required|string|max:255",
		"icon": "nullable|string|max:255",
	})
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "Validation failed"})
	}
	if validator.Fails() {
		return ctx.Response().Status(http.StatusUnprocessableEntity).Json(http.Json{"message": "Validation failed", "errors": validator.Errors().All()})
	}

	facility, err := c.service.Update(id, request)
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"data": facility})
}

func (c *FacilityController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	if err := c.service.Delete(id); err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"message": "Facility deleted successfully"})
}
