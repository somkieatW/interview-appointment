package appointment

import (
	"context"
	"github.com/somkieatW/interview-appointment/pkg/models"
)

type AppointmentUseCase interface {
	List(ctx context.Context, request *models.AppointmentListRequest) (*models.AppointmentListResponse, error)
	Info(ctx context.Context, request *models.AppointmentInfoRequest) (*models.AppointmentInfoResponse, error)
	Update(ctx context.Context, request *models.AppointmentUpdateRequest) (*models.AppointmentUpdateResponse, error)
	Archive(ctx context.Context, request *models.AppointmentArchiveRequest) (*models.AppointmentArchiveResponse, error)
}
