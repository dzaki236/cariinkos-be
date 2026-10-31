package dto

type UpdateProfileRequest struct {
	Name                    string  `json:"name"`
	Email                   string  `json:"email"`
	OldPassword             *string `json:"old_password,omitempty"`
	NewPassword             *string `json:"new_password,omitempty"`
	NewPasswordConfirmation *string `json:"new_password_confirmation,omitempty"`
}
