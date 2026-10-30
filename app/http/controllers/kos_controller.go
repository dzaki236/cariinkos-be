package controllers

import (
	"cariinkos/app/dto"
	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	serviceContracts "cariinkos/app/services/contracts"

	"github.com/goravel/framework/contracts/http"
)

type KosController struct {
	service serviceContracts.KosService
}

func NewKosController(service serviceContracts.KosService) *KosController {
	return &KosController{service: service}
}

func (c *KosController) Index(ctx http.Context) http.Response {
	kos, err := c.service.GetAllData()
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "Failed to retrieve kos"})
	}
	return ctx.Response().Success().Json(http.Json{"data": kos})
}

func (c *KosController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	kos, err := c.service.GetByID(id)
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"data": kos})
}

func (c *KosController) Store(ctx http.Context) http.Response {
	var request dto.CreateKosRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return ctx.Response().Status(http.DefaultAbortStatus).Json(http.Json{"message": "Invalid request body"})
	}
	validator, err := ctx.Request().Validate(map[string]any{
		"name":          "required|string|max:255",
		"address":       "required|string",
		"province":      "required|string|max:255",
		"city":          "required|string|max:255",
		"district":      "required|string|max:255",
		"price_monthly": "required|numeric|min:1",
		"type":          "in:mixed,man,woman",
		"status":        "in:draft,published,archived",
	})
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "Validation failed"})
	}
	if validator.Fails() {
		return ctx.Response().Status(http.StatusUnprocessableEntity).Json(http.Json{"message": "Validation failed", "errors": validator.Errors().All()})
	}

	userID, err := facades.Auth(ctx).ID()
	if err != nil || userID == "" {
		return ctx.Response().Status(http.StatusUnauthorized).Json(http.Json{"message": "Unauthorized"})
	}

	kos, err := c.service.Create(userID, request)
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Status(http.StatusCreated).Json(http.Json{"data": kos})
}

func (c *KosController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var request dto.UpdateKosRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return ctx.Response().Status(http.StatusNotFound).Json(http.Json{"message": "Invalid request body"})
	}
	validator, err := ctx.Request().Validate(map[string]any{
		"name":          "required|string|max:255",
		"address":       "required|string",
		"province":      "required|string|max:255",
		"city":          "required|string|max:255",
		"district":      "required|string|max:255",
		"price_monthly": "required|numeric|min:1",
		"type":          "required|in:mixed,man,woman",
		"status":        "required|in:draft,published,archived",
	})
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "Validation failed"})
	}
	if validator.Fails() {
		return ctx.Response().Status(http.StatusUnprocessableEntity).Json(http.Json{"message": "Validation failed", "errors": validator.Errors().All()})
	}

	kos, err := c.service.Update(id, request)
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"data": kos})
}

func (c *KosController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	if err := c.service.Delete(id); err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"message": "Kos deleted successfully"})
}
