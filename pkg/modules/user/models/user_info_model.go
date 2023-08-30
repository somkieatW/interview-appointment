package models

import (
	"github.com/somkieatW/candidate/pkg/domain"
)

type UserInfoRequest struct {
	ID string `json:"id"`
}

type UserInfoResponse struct {
	Data    *domain.User `json:"data"`
	Success bool         `json:"success"`
}
