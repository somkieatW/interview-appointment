package usecase

import (
	"github.com/somkieatW/interview-appointment/pkg/core/registry"
	"github.com/somkieatW/interview-appointment/pkg/core/registry/core"
	"github.com/somkieatW/interview-appointment/pkg/modules/appointment"
)

type appointmentUseCase struct {
	CoreRegistry       *core.CoreRegistry
	RepositoryRegistry *registry.RepositoryRegistry
}

func NewAppointmentUseCase(coreRegistry *core.CoreRegistry, repositoryRegistry *registry.RepositoryRegistry) appointment.AppointmentUseCase {
	return &appointmentUseCase{
		CoreRegistry:       coreRegistry,
		RepositoryRegistry: repositoryRegistry,
	}
}
