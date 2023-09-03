package mock

import (
	"context"
	"github.com/somkieatW/interview-appointment/pkg/domain"
	"github.com/somkieatW/interview-appointment/pkg/models"
	"github.com/stretchr/testify/mock"
)

// MockAppointmentRepository is a mock implementation of the AppointmentRepository interface for testing purposes.
type MockAppointmentRepository struct {
	mock.Mock
	ListFunc   func(ctx context.Context, request *models.AppointmentListRequest) (*models.AppointmentListResponse, error)
	UpdateFunc func(ctx context.Context, appointment *domain.Appointment) error
}

func (m *MockAppointmentRepository) IsExisted(ctx context.Context, id string) bool {
	args := m.Called(ctx, id)
	return args.Bool(0)
}

func (m *MockAppointmentRepository) IsNotExisted(ctx context.Context, id string) bool {
	args := m.Called(ctx, id)
	return args.Bool(0)
}

func (m *MockAppointmentRepository) List(ctx context.Context, obj *models.AppointmentListRequest) (*[]models.AppointmentListData, error) {
	args := m.Called(ctx, obj)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	data := args.Get(0).(*[]models.AppointmentListData)
	return data, args.Error(1)
}

func (m *MockAppointmentRepository) Info(ctx context.Context, obj *models.AppointmentInfoRequest) (*models.AppointmentInfoData, error) {
	args := m.Called(ctx, obj)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	data := args.Get(0).(*models.AppointmentInfoData)
	return data, args.Error(1)
}

func (m *MockAppointmentRepository) Update(ctx context.Context, obj *domain.Appointment) error {
	args := m.Called(ctx, obj)
	return args.Error(0)
}
