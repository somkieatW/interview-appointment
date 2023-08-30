package models

import (
	"github.com/somkieatW/candidate/pkg/domain"
)

type AppointmentListRequest struct {
	Offset   int `json:"offset"  example:"1"`
	PageSize int `json:"pageSize" example:"10"`
}

type AppointmentListResponse struct {
	Data    *[]domain.Appointment `json:"data"`
	Success bool                  `json:"success"`
}
