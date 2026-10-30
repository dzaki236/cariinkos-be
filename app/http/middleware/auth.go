package middleware

import (
	"strings"

	"cariinkos/app/facades"

	"github.com/goravel/framework/contracts/http"
)

type Auth struct{}

func NewHandleAuth() *Auth {
	return &Auth{}
}

func (a *Auth) Handle(ctx http.Context) {
	parts := strings.Fields(ctx.Request().Header("Authorization"))

	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		ctx.Response().Json(http.StatusUnauthorized, map[string]any{
			"message": "Unauthorized",
			"error":   "Authorization header must use Bearer token format",
		}).Abort()

		return
	}

	// Parse dan validasi JWT.
	_, err := facades.Auth(ctx).Parse(parts[1])

	if err != nil {
		ctx.Response().Json(http.StatusUnauthorized, map[string]any{
			"message": "Unauthorized",
			"error":   "Invalid or expired token",
		}).Abort()

		return
	}

	// Lanjutkan ke middleware/controller berikutnya.
	ctx.Request().Next()
}

func (a *Auth) Signature() string {
	return "auth"
}
