package middleware

import (
	"cariinkos/app/facades"
	"cariinkos/app/repositories"

	"github.com/goravel/framework/contracts/http"
)

type Role struct{ role string }

func NewRoleAuth(role string) *Role { return &Role{role: role} }

func (r *Role) Handle(ctx http.Context) {
	userID, err := facades.Auth(ctx).ID()
	if err != nil || userID == "" {
		ctx.Response().Json(http.StatusUnauthorized, map[string]any{"message": "Unauthorized"}).Abort()
		return
	}
	user, err := repositories.NewUserAuthRepository().FindByID(userID)
	if err != nil || user.Role != r.role {
		ctx.Response().Json(http.StatusForbidden, map[string]any{"message": "Forbidden"}).Abort()
		return
	}
	ctx.Request().Next()
}

func (r *Role) Signature() string { return "role:" + r.role }
