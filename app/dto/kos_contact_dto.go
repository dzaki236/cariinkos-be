package dto

type CreateKosContactRequest struct {
	ContactName string  `json:"contact_name"`
	Phone       string  `json:"phone"`
	Whatsapp    *string `json:"whatsapp"`
}

type UpdateKosContactRequest struct {
	ContactName string  `json:"contact_name"`
	Phone       string  `json:"phone"`
	Whatsapp    *string `json:"whatsapp"`
}
