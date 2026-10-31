package controllers

import (
	"cariinkos/app/dto"
	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	serviceContracts "cariinkos/app/services/contracts"

	"github.com/goravel/framework/contracts/http"
)

type KosFavoriteController struct {
	service serviceContracts.KosFavoriteService
}

func NewKosFavoriteController(service serviceContracts.KosFavoriteService) *KosFavoriteController {
	return &KosFavoriteController{service: service}
}
func (c *KosFavoriteController) userID(ctx http.Context) (string, error) {
	return facades.Auth(ctx).ID()
}
func (c *KosFavoriteController) Index(ctx http.Context) http.Response {
	userID, err := c.userID(ctx)
	if err != nil {
		return ctx.Response().Status(http.StatusUnauthorized).Json(http.Json{"message": "Unauthorized"})
	}
	data, err := c.service.GetByUser(userID)
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"data": data})
}
func (c *KosFavoriteController) Store(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	userID, err := c.userID(ctx)
	if err != nil {
		return ctx.Response().Status(http.StatusUnauthorized).Json(http.Json{"message": "Unauthorized"})
	}
	var request dto.UpdateFavoriteRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return ctx.Response().Status(http.StatusBadRequest).Json(http.Json{"message": "Invalid request body"})
	}
	data, err := c.service.Create(userID, id, request)
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Status(http.StatusCreated).Json(http.Json{"data": data})
}
func (c *KosFavoriteController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	userID, err := c.userID(ctx)
	if err != nil {
		return ctx.Response().Status(http.StatusUnauthorized).Json(http.Json{"message": "Unauthorized"})
	}
	var request dto.UpdateFavoriteRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return ctx.Response().Status(http.StatusBadRequest).Json(http.Json{"message": "Invalid request body"})
	}
	data, err := c.service.Update(userID, id, request)
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"data": data})
}
func (c *KosFavoriteController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	userID, err := c.userID(ctx)
	if err != nil {
		return ctx.Response().Status(http.StatusUnauthorized).Json(http.Json{"message": "Unauthorized"})
	}
	if err := c.service.Delete(userID, id); err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"message": "Favorite deleted successfully"})
}
