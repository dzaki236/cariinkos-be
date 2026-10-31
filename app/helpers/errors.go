package helpers

import (
	"errors"

	"github.com/goravel/framework/contracts/http"
)

var ErrDataNotFound = errors.New("Data not found")

// var ErrFavoriteAlreadyExists = errors.New("Favorite already exists")
var ErrDataAlreadyExists = errors.New("Data already exists")

// var ErrReviewAlreadyExists = errors.New("Review already exists")
var ErrPhotoRequired = errors.New("photo file is required")
var ErrInvalidCredentials = errors.New("invalid email or password")
var ErrEmailAlreadyRegistered = errors.New("email is already registered")
var ErrInvalidProfilePassword = errors.New("invalid profile password")
var ErrPasswordConfirmationMismatch = errors.New("new password confirmation does not match")
var ErrNotManageableUser = errors.New("only users with role user can be managed")

func ErrorResponse(ctx http.Context, err error) http.Response {
	if errors.Is(err, ErrDataNotFound) {
		return ctx.Response().Status(http.StatusNotFound).Json(http.Json{"message": "Data not found", "unrealmsg": err.Error()})
	}

	if errors.Is(err, ErrEmailAlreadyRegistered) {
		return ctx.Response().Status(http.StatusConflict).Json(http.Json{"message": "Email is already registered", "unrealmsg": err.Error()})
	}

	if errors.Is(err, ErrDataAlreadyExists) {
		return ctx.Response().Status(http.StatusConflict).Json(http.Json{"message": "Data already exists", "unrealmsg": err.Error()})
	}
	
	return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "An unexpected error occurred", "unrealmsg": err.Error()})
}
