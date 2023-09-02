package usecase

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/somkieatW/candidate/pkg/modules/appointment/models"
)

func infoValidate(request *models.AppointmentInfoRequest) error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.ID, validation.Required, is.UUID),
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *appointmentUseCase) Info(ctx context.Context, request *models.AppointmentInfoRequest) (*models.AppointmentInfoResponse, error) {
	if err := infoValidate(request); err != nil {
		return nil, err
	}

	data, err := u.RepositoryRegistry.AppointmentRepository.Info(ctx, request)
	if err != nil {
		return nil, err
	}

	res := &models.AppointmentInfoResponse{
		Success: true,
		Data:    data,
	}

	return res, nil
}
