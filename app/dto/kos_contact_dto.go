package dto

type CreateKosContactRequest struct {
	Phone    string  `json:"phone"`
	Whatsapp *string `json:"whatsapp"`
}

type UpdateKosContactRequest struct {
	Phone    string  `json:"phone"`
	Whatsapp *string `json:"whatsapp"`
}
