package models

import (
	"github.com/somkieatW/interview-appointment/pkg/domain"
)

type AppointmentListRequest struct {
	Offset   int `json:"offset"  example:"1"`
	PageSize int `json:"pageSize" example:"10"`
}

type AppointmentListResponse struct {
	Data    *[]AppointmentListData `json:"data"`
	Success bool                   `json:"success"`
}

type AppointmentListData struct {
	*domain.Appointment
	DisplayName string `json:"displayName"`
}
