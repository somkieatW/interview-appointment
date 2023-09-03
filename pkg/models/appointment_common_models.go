package models

type Appointment struct {
	ID          string `json:"id"`
	Topic       string `json:"topic"`
	State       string `json:"state"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	CreatedBy   string `json:"createdBy"`
}
