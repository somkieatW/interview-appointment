package usecase

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/somkieatW/interview-appointment/pkg/core/constants"
	"github.com/somkieatW/interview-appointment/pkg/domain"
	"github.com/somkieatW/interview-appointment/pkg/modules/appointment/models"
	"time"
)

func updateValidate(request *models.AppointmentUpdateRequest) error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.ID, validation.Required, is.UUID),
		validation.Field(&request.State, validation.Required, validation.In(
			constants.StateTodo, constants.StateInProgress, constants.StateDone,
		)),
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *appointmentUseCase) Update(ctx context.Context, request *models.AppointmentUpdateRequest) (*models.AppointmentUpdateResponse, error) {
	if err := updateValidate(request); err != nil {
		return nil, err
	}

	obj := &domain.Appointment{
		ID:          request.ID,
		State:       request.State,
		UpdatedDate: time.Now(),
	}

	err := u.RepositoryRegistry.AppointmentRepository.Update(ctx, obj)
	if err != nil {
		return nil, err
	}

	res := &models.AppointmentUpdateResponse{
		Success: true,
	}

	return res, nil
}
