package contracts

import (
	"errors"

	"cariinkos/app/dto"
	"cariinkos/app/models"

	"github.com/goravel/framework/contracts/http"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

type UserAuthService interface {
	Login(ctx http.Context, request dto.LoginRequest) (models.User, string, error)
	Logout(ctx http.Context) error
}
