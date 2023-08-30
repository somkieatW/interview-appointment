package models

import (
	"github.com/somkieatW/candidate/pkg/domain"
)

type AppointmentInfoRequest struct {
	ID string `json:"id"`
}

type AppointmentInfoResponse struct {
	Data    *domain.Appointment `json:"data"`
	Success bool                `json:"success"`
}
