package models

import (
	"github.com/somkieatW/candidate/pkg/domain"
)

type CommentCreateRequest struct {
	AppointmentID string `json:"appointmentId"`
	UserID        string `json:"userId"`
	Content       string `json:"content"`
}

type CommentCreateResponse struct {
	Data    *domain.Comment `json:"data"`
	Success bool            `json:"success"`
}
