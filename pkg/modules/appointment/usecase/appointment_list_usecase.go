package usecase

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/somkieatW/candidate/pkg/modules/appointment/models"
)

func listValidate(request *models.AppointmentListRequest) error {
	err := validation.ValidateStruct(request)

	if err != nil {
		return err
	}

	return nil
}

func (u *candidateUseCase) List(ctx context.Context, request *models.AppointmentListRequest) (*models.AppointmentListResponse, error) {
	if err := listValidate(request); err != nil {
		return nil, err
	}

	data, err := u.RepositoryRegistry.AppointmentRepository.List(ctx, request.AppointmentListCriteria)
	if err != nil {
		return nil, err
	}

	res := &models.AppointmentListResponse{
		Success: true,
		Data:    data,
	}

	return res, nil
}