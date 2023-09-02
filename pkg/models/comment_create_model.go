package models

type CommentCreateRequest struct {
	AppointmentID string `json:"appointmentId"`
	UserID        string `json:"userId"`
	Content       string `json:"content"`
}

type CommentCreateResponse struct {
	Success bool `json:"success"`
}
