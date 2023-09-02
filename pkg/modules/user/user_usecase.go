package user

import (
	"context"
	"github.com/somkieatW/interview-appointment/pkg/modules/user/models"
)

type UserUseCase interface {
	Info(ctx context.Context, request *models.UserInfoRequest) (*models.UserInfoResponse, error)
}
