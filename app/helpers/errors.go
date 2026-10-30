package helpers

import (
	"errors"

	"github.com/goravel/framework/contracts/http"
)

var ErrDataNotFound = errors.New("Data not found")
var ErrPhotoRequired = errors.New("photo file is required")

func ErrorResponse(ctx http.Context, err error) http.Response {
	if errors.Is(err, ErrDataNotFound) {
		return ctx.Response().Status(http.StatusNotFound).Json(http.Json{"message": "Data not found"})
	}
	return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"message": "An unexpected error occurred", "unrealmsg": err.Error()})
}
