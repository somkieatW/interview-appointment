package appointment

import (
	"context"
	"github.com/somkieatW/interview-appointment/pkg/modules/appointment/models"
)

type AppointmentUseCase interface {
	List(ctx context.Context, request *models.AppointmentListRequest) (*models.AppointmentListResponse, error)
	Info(ctx context.Context, request *models.AppointmentInfoRequest) (*models.AppointmentInfoResponse, error)
}
