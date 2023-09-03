package usecase

import (
	"context"
	"github.com/somkieatW/interview-appointment/pkg/core/registry"
	"github.com/somkieatW/interview-appointment/pkg/modules/appointment/mock"
	"testing"

	"github.com/somkieatW/interview-appointment/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	mockRepo := &mock.MockAppointmentRepository{}

	useCase := &appointmentUseCase{
		RepositoryRegistry: &registry.RepositoryRegistry{
			AppointmentRepository: mockRepo,
		},
	}

	// Test case 1: Valid request
	validRequest := &models.AppointmentListRequest{}
	mockData := []models.AppointmentListData{{
		Appointment: &models.Appointment{
			ID:          "id",
			Topic:       "สัมพาทงาน",
			State:       "To Do",
			Description: "Test",
		},
		DisplayName: "โรบินฮู้ด",
	}}

	mockRepo.On("List", context.Background(), validRequest).Return(&mockData, nil)

	response, err := useCase.List(context.Background(), validRequest)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.True(t, response.Success)
	assert.Equal(t, &mockData, response.Data)

	mockRepo.AssertCalled(t, "List", context.Background(), validRequest)

	// Test case 2: Error during validation
	invalidRequest := &models.AppointmentListRequest{
		Offset:   -1,
		PageSize: 20,
	}

	response, err = useCase.List(context.Background(), invalidRequest)

	assert.Error(t, err)
	assert.Nil(t, response)

	response, err = useCase.List(context.Background(), invalidRequest)

	assert.Error(t, err)
	assert.Nil(t, response)
}
