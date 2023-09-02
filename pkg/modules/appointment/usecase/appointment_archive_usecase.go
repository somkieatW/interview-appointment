package usecase

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/somkieatW/interview-appointment/pkg/core/constants"
	"github.com/somkieatW/interview-appointment/pkg/domain"
	"github.com/somkieatW/interview-appointment/pkg/models"
	"time"
)

func archiveValidate(request *models.AppointmentArchiveRequest) error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.ID, validation.Required, is.UUID),
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *appointmentUseCase) Archive(ctx context.Context, request *models.AppointmentArchiveRequest) (*models.AppointmentArchiveResponse, error) {
	if err := archiveValidate(request); err != nil {
		return nil, err
	}

	obj := &domain.Appointment{
		ID:        request.ID,
		Status:    constants.StatusInactive,
		UpdatedAt: time.Now(),
	}

	err := u.RepositoryRegistry.AppointmentRepository.Update(ctx, obj)
	if err != nil {
		return nil, err
	}

	res := &models.AppointmentArchiveResponse{
		Success: true,
	}

	return res, nil
}
