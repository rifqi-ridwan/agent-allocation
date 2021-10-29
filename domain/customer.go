package domain

type Customer struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	IsResolved bool   `json:"is_resolved"`
	RoomID     string `json:"room_id"`
}
