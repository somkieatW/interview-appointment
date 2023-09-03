package models

type AppointmentListRequest struct {
	Offset   int `json:"offset"  example:"1"`
	PageSize int `json:"pageSize" example:"10"`
}

type AppointmentListResponse struct {
	Data    *[]AppointmentListData `json:"data"`
	Success bool                   `json:"success"`
}

type AppointmentListData struct {
	*Appointment
	DisplayName string `json:"displayName"`
}
