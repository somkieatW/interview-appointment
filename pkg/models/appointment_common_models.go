package models

type Appointment struct {
	Topic       string `json:"topic"`
	State       string `json:"state"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	CreatedBy   string `json:"createdBy"`
}
