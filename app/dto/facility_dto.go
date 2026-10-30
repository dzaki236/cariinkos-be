package dto

type CreateFacilityRequest struct {
	Name string  `json:"name"`
	Icon *string `json:"icon"`
}
type UpdateFacilityRequest struct {
	Name string  `json:"name"`
	Icon *string `json:"icon"`
}
