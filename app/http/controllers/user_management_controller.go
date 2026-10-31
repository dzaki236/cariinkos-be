package controllers

import (
	"errors"

	"cariinkos/app/dto"
	"cariinkos/app/helpers"
	serviceContracts "cariinkos/app/services/contracts"

	"github.com/goravel/framework/contracts/http"
)

type UserManagementController struct {
	service serviceContracts.UserManagementService
}

func NewUserManagementController(service serviceContracts.UserManagementService) *UserManagementController {
	return &UserManagementController{service: service}
}

func (c *UserManagementController) Index(ctx http.Context) http.Response {
	users, err := c.service.GetUsers()
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"data": users})
}

func (c *UserManagementController) UpdateStatus(ctx http.Context) http.Response {
	var request dto.UpdateUserStatusRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return ctx.Response().Status(http.StatusBadRequest).Json(http.Json{"message": "Invalid request body"})
	}
	user, err := c.service.UpdateUserStatus(ctx.Request().Route("id"), request.Active)
	if errors.Is(err, serviceContracts.ErrNotManageableUser) {
		return ctx.Response().Status(http.StatusUnprocessableEntity).Json(http.Json{"message": err.Error()})
	}
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"data": user})
}
