package domain

type Agent struct {
	ID                   int  `json:"id"`
	IsAvailable          bool `json:"is_available"`
	CurrentCustomerCount int  `json:"current_customer_count"`
}

type AllAgent struct {
	Data struct {
		Agent struct {
			Data []Agent `json:"data"`
		} `json:"agents"`
	} `json:"data"`
}

type QueuePayload struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	IsResolved bool   `json:"is_resolved"`
	RoomID     string `json:"room_id"`
}
