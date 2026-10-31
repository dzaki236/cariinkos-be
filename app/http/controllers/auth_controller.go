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
		if errors.Is(err, helpers.ErrInvalidCredentials) {
			return ctx.Response().Status(http.StatusUnauthorized).Json(http.Json{"message": "Invalid email or password"})
		}
		return helpers.ErrorResponse(ctx, err)
	}

	return ctx.Response().Success().Json(http.Json{"message": "Login successful", "token": token, "user": user})
}

func (c *AuthController) Register(ctx http.Context) http.Response {
	var request dto.RegisterRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return ctx.Response().Status(http.StatusBadRequest).Json(http.Json{"message": "Invalid request body"})
	}

	validator, err := ctx.Request().Validate(map[string]any{
		"name":     "required|string|max:255",
		"email":    "required|email|max:255",
		"password": "required|string|min:8",
	})
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": err.Error()})
	}
	if validator.Fails() {
		return ctx.Response().Status(http.StatusUnprocessableEntity).Json(http.Json{"message": "Validation failed", "errors": validator.Errors().All()})
	}

	user, err := c.service.Register(request)
	if errors.Is(err, helpers.ErrDataAlreadyExists) {
		return ctx.Response().Status(http.StatusConflict).Json(http.Json{"message": err.Error()})
	}
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Status(http.StatusCreated).Json(http.Json{"message": "Registration successful", "user": user})
}

func (c *AuthController) Logout(ctx http.Context) http.Response {
	if err := c.service.Logout(ctx); err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "Logout failed"})
	}
	return ctx.Response().Success().Json(http.Json{"message": "Logout successful"})
}

func (c *AuthController) Profile(ctx http.Context) http.Response {
	var request dto.UpdateProfileRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return ctx.Response().Status(http.StatusBadRequest).Json(http.Json{"message": "Invalid request body"})
	}

	validator, err := ctx.Request().Validate(map[string]any{
		"name":                      "required|string|max:255",
		"email":                     "required|email|max:255",
		"old_password":              "required_with:new_password|string",
		"new_password":              "nullable|string|min:8",
		"new_password_confirmation": "required_with:new_password|same:new_password",
	})
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": err.Error()})
	}
	if validator.Fails() {
		return ctx.Response().Status(http.StatusUnprocessableEntity).Json(http.Json{"message": "Validation failed", "errors": validator.Errors().All()})
	}

	user, err := c.service.UpdateProfile(ctx, request)
	if err != nil {
		return helpers.ErrorResponse(ctx, err)
	}
	return ctx.Response().Success().Json(http.Json{"message": "Profile updated successfully", "user": user})
}
