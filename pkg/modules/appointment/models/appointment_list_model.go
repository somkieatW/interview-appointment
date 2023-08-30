package models

import (
	"github.com/somkieatW/candidate/pkg/domain"
	"github.com/somkieatW/candidate/pkg/repository/repomodels"
)

type AppointmentListRequest struct {
	*repomodels.AppointmentListCriteria
}

type AppointmentListResponse struct {
	Data    *[]domain.Appointment `json:"data"`
	Success bool                  `json:"success"`
}
