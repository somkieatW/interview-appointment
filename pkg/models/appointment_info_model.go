package models

type AppointmentInfoRequest struct {
	ID string `json:"id"`
}

type AppointmentInfoResponse struct {
	Data    *AppointmentInfoData `json:"data"`
	Success bool                 `json:"success"`
}

type AppointmentInfoData struct {
	*Appointment
	Comments *[]Comments `json:"comments"`
}
