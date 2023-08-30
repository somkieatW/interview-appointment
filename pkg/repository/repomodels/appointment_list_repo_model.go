package repomodels

import "github.com/somkieatW/candidate/pkg/core/models"

type AppointmentListCriteria struct {
	Limit *models.PageLimit `json:"limit"`
}
