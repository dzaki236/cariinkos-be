package contracts

import (
	"errors"

	"cariinkos/app/dto"
	"cariinkos/app/models"

	"github.com/goravel/framework/contracts/http"
)

var ErrInvalidCredentials = errors.New("invalid email or password")
var ErrEmailAlreadyRegistered = errors.New("email is already registered")

type UserAuthService interface {
	Login(ctx http.Context, request dto.LoginRequest) (models.User, string, error)
	Register(request dto.RegisterRequest) (models.User, error)
	Logout(ctx http.Context) error
	UpdateProfile(ctx http.Context, request dto.UpdateProfileRequest) (models.User, error)
}
