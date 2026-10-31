package controllers

import (
	"cariinkos/app/dto"
	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	serviceContracts "cariinkos/app/services/contracts"

	"github.com/goravel/framework/contracts/http"
)

type KosReviewController struct {
	service serviceContracts.KosReviewService
}

func NewKosReviewController(service serviceContracts.KosReviewService) *KosReviewController {
	return &KosReviewController{service: service}
}
func (c *KosReviewController) Index(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	data, err := c.service.GetByKos(id)
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"data": data})
}
func (c *KosReviewController) Store(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	userID, err := facades.Auth(ctx).ID()
	if err != nil {
		return ctx.Response().Status(http.StatusUnauthorized).Json(http.Json{"message": "Unauthorized"})
	}
	var request dto.ReviewRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return ctx.Response().Status(http.StatusBadRequest).Json(http.Json{"message": "Invalid request body"})
	}
	validator, err := ctx.Request().Validate(map[string]any{"rating": "required|numeric|min:1|max:5", "comment": "required|string|max:2000"})
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "Validation failed"})
	}
	if validator.Fails() {
		return ctx.Response().Status(http.StatusUnprocessableEntity).Json(http.Json{"message": "Validation failed", "errors": validator.Errors().All()})
	}
	data, err := c.service.Create(userID, id, request)
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Status(http.StatusCreated).Json(http.Json{"data": data})
}
func (c *KosReviewController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	userID, err := facades.Auth(ctx).ID()
	if err != nil {
		return ctx.Response().Status(http.StatusUnauthorized).Json(http.Json{"message": "Unauthorized"})
	}
	var request dto.ReviewRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return ctx.Response().Status(http.StatusBadRequest).Json(http.Json{"message": "Invalid request body"})
	}
	data, err := c.service.Update(userID, id, request)
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"data": data})
}
func (c *KosReviewController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	userID, err := facades.Auth(ctx).ID()
	if err != nil {
		return ctx.Response().Status(http.StatusUnauthorized).Json(http.Json{"message": "Unauthorized"})
	}
	if err := c.service.Delete(userID, id); err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"message": "Review deleted successfully"})
}
