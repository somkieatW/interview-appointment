package usecase

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/somkieatW/candidate/pkg/domain"
	"github.com/somkieatW/candidate/pkg/modules/comment/models"
	"time"
)

func (u *commentUseCase) createValidate(request *models.CommentCreateRequest) error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.AppointmentID, validation.Required, is.UUID),
		validation.Field(&request.UserID, validation.Required, is.UUID),
		validation.Field(&request.Content, validation.Required),
	)

	if err != nil {
		return err
	}

	return nil
}

func (u *commentUseCase) Create(ctx context.Context, request *models.CommentCreateRequest) (*models.CommentCreateResponse, error) {
	if err := u.createValidate(request); err != nil {
		return nil, err
	}

	if u.RepositoryRegistry.UserRepository.IsNotExisted(ctx, request.UserID) {
		return nil, errors.New("This User is Not existed")
	}

	if u.RepositoryRegistry.AppointmentRepository.IsNotExisted(ctx, request.AppointmentID) {
		return nil, errors.New("This Appointment is not existed")
	}

	obj := &domain.Comment{
		ID:            uuid.NewString(),
		AppointmentID: request.AppointmentID,
		UserID:        request.UserID,
		Content:       request.Content,
		CreatedDate:   time.Now(),
	}

	data, err := u.RepositoryRegistry.CommentRepository.Create(ctx, obj)
	if err != nil {
		return nil, err
	}

	res := &models.CommentCreateResponse{
		Success: true,
		Data:    data,
	}

	return res, nil
}
