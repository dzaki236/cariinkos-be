package controllers

import (
	"errors"

	"cariinkos/app/dto"
	"cariinkos/app/helpers"
	serviceContracts "cariinkos/app/services/contracts"

	"github.com/goravel/framework/contracts/http"
)

type AuthController struct {
	service serviceContracts.UserAuthService
}

func NewAuthController(service serviceContracts.UserAuthService) *AuthController {
	return &AuthController{service: service}
}

func (c *AuthController) Login(ctx http.Context) http.Response {
	var request dto.LoginRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return ctx.Response().Status(http.StatusBadRequest).Json(http.Json{"message": "Invalid request body"})
	}

	validator, err := ctx.Request().Validate(map[string]any{
		"email":    "required|email",
		"password": "required|string",
	})
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": err})
	}
	if validator.Fails() {
		return ctx.Response().Status(http.StatusUnprocessableEntity).Json(http.Json{"message": "Validation failed", "errors": validator.Errors().All()})
	}

	user, token, err := c.service.Login(ctx, request)
	if err != nil {
		if errors.Is(err, serviceContracts.ErrInvalidCredentials) {
			return ctx.Response().Status(http.StatusUnauthorized).Json(http.Json{"message": "Invalid email or password"})
		}
		return helpers.ErrorResponse(ctx, err)
	}

	return ctx.Response().Success().Json(http.Json{"message": "Login successful", "token": token, "user": user})
}

func (c *AuthController) Logout(ctx http.Context) http.Response {
	if err := c.service.Logout(ctx); err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "Logout failed"})
	}
	return ctx.Response().Success().Json(http.Json{"message": "Logout successful"})
}
