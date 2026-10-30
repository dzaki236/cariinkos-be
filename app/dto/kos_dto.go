package dto

type CreateKosRequest struct {
	Name              string  `json:"name"`
	Description       *string `json:"description"`
	Address           string  `json:"address"`
	Province          string  `json:"province"`
	City              string  `json:"city"`
	District          string  `json:"district"`
	Subdistrict       *string `json:"subdistrict"`
	GmapsLink         *string `json:"gmaps_link"`
	PriceMonthly      uint64  `json:"price_monthly"`
	Type              string  `json:"type"`
	Status            string  `json:"status"`
	RoomSize          *uint64 `json:"room_size"`
	TotalRooms        uint64  `json:"total_rooms"`
	AvailableRooms    uint64  `json:"available_rooms"`
	MinimumRentMonths uint64  `json:"minimum_rent_months"`
	Rules             *string `json:"rules"`
}

type UpdateKosRequest struct {
	Name              string  `json:"name"`
	Description       *string `json:"description"`
	Address           string  `json:"address"`
	Province          string  `json:"province"`
	City              string  `json:"city"`
	District          string  `json:"district"`
	Subdistrict       *string `json:"subdistrict"`
	GmapsLink         *string `json:"gmaps_link"`
	PriceMonthly      uint64  `json:"price_monthly"`
	Type              string  `json:"type"`
	Status            string  `json:"status"`
	RoomSize          *uint64 `json:"room_size"`
	TotalRooms        uint64  `json:"total_rooms"`
	AvailableRooms    uint64  `json:"available_rooms"`
	MinimumRentMonths uint64  `json:"minimum_rent_months"`
	Rules             *string `json:"rules"`
}
