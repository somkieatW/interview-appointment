package usecase

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/somkieatW/interview-appointment/pkg/modules/user/models"
)

func infoValidate(request *models.UserInfoRequest) error {
	err := validation.ValidateStruct(request)

	if err != nil {
		return err
	}

	return nil
}

func (u *userUseCase) Info(ctx context.Context, request *models.UserInfoRequest) (*models.UserInfoResponse, error) {
	if err := infoValidate(request); err != nil {
		return nil, err
	}

	data, err := u.RepositoryRegistry.UserRepository.Info(ctx, request)
	if err != nil {
		return nil, err
	}

	res := &models.UserInfoResponse{
		Success: true,
		Data:    data,
	}

	return res, nil
}
