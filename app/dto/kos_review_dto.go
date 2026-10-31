package dto

type ReviewRequest struct {
	Rating  uint8  `json:"rating"`
	Comment string `json:"comment"`
}
