package appointment

import (
	"context"
	models2 "github.com/somkieatW/interview-appointment/pkg/models"
)

type AppointmentUseCase interface {
	List(ctx context.Context, request *models2.AppointmentListRequest) (*models2.AppointmentListResponse, error)
	Info(ctx context.Context, request *models2.AppointmentInfoRequest) (*models2.AppointmentInfoResponse, error)
}
