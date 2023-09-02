package models

import (
	"github.com/somkieatW/candidate/pkg/domain"
	"time"
)

type AppointmentInfoRequest struct {
	ID string `json:"id"`
}

type AppointmentInfoResponse struct {
	Data    *AppointmentInfoData
	Success bool `json:"success"`
}

type AppointmentInfoData struct {
	Topic       string            `json:"topic"`
	Description string            `json:"description"`
	State       string            `json:"state"`
	CreatedBy   string            `json:"createdBy"`
	CreatedDate time.Time         `json:"createdDate"`
	DisplayName string            `json:"displayName"`
	Comments    *[]domain.Comment `json:"comments"`
}
