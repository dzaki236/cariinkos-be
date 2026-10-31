package dto

type UpdateProfileRequest struct {
	Name                    string  `json:"name"`
	Email                   string  `json:"email"`
	OldPassword             *string ``
	NewPassword             *string ``
	NewPasswordConfirmation *string ``
}
