package usecase

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/somkieatW/interview-appointment/pkg/models"
)

func listValidate(request *models.AppointmentListRequest) error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.Offset, is.Digit, is.Digit, validation.Min(0)),
		validation.Field(&request.PageSize, is.Digit, validation.Min(0)),
	)

	if err != nil {
		return err
	}

	return nil
}

func (u *appointmentUseCase) List(ctx context.Context, request *models.AppointmentListRequest) (*models.AppointmentListResponse, error) {
	if err := listValidate(request); err != nil {
		return nil, err
	}

	data, err := u.RepositoryRegistry.AppointmentRepository.List(ctx, request)
	if err != nil {
		return nil, err
	}

	res := &models.AppointmentListResponse{
		Success: true,
		Data:    data,
	}

	return res, nil
}
