package models

type AppointmentArchiveRequest struct {
	ID string `json:"id"`
}

type AppointmentArchiveResponse struct {
	Success bool `json:"success"`
}
