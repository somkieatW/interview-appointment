package usecase

import (
	"context"
	"errors"
	"github.com/somkieatW/interview-appointment/pkg/core/registry"
	"github.com/somkieatW/interview-appointment/pkg/domain"
	"github.com/somkieatW/interview-appointment/pkg/models"
	"github.com/somkieatW/interview-appointment/pkg/modules/appointment/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArchive(t *testing.T) {
	// Create a sample AppointmentArchiveRequest
	request := &models.AppointmentArchiveRequest{
		ID: "a2e223eb-5141-4d74-84f1-db85c3aa5a09",
	}

	mockRepo := &mock.MockAppointmentRepository{
		UpdateFunc: func(ctx context.Context, appointment *domain.Appointment) error {
			return nil
		},
	}

	// Create an instance of the appointment use case with the mock repository
	useCase := &appointmentUseCase{
		RepositoryRegistry: &registry.RepositoryRegistry{
			AppointmentRepository: mockRepo,
		},
	}

	// Test case 1: Successful archive
	t.Run("Successful Archive", func(t *testing.T) {
		response, err := useCase.Archive(context.Background(), request)

		assert.NoError(t, err, "Expected no error")
		assert.NotNil(t, response, "Expected a non-nil response")
		assert.True(t, response.Success, "Expected a successful response")
	})

	// Test case 2: Error in updating the appointment
	t.Run("Error in Update", func(t *testing.T) {
		mockRepo.UpdateFunc = func(ctx context.Context, appointment *domain.Appointment) error {
			// Simulate an error in updating the appointment
			return errors.New("update failed")
		}

		response, err := useCase.Archive(context.Background(), request)

		assert.Error(t, err, "Expected an error")
		assert.Nil(t, response, "Expected a nil response")
	})
}
