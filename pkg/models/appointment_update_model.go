package models

type AppointmentUpdateRequest struct {
	ID    string `json:"id"`
	State string `json:"state"`
}

type AppointmentUpdateResponse struct {
	Success bool `json:"success"`
}
