package usecase

import (
	"github.com/somkieatW/candidate/pkg/core/registry"
	"github.com/somkieatW/candidate/pkg/core/registry/core"
	"github.com/somkieatW/candidate/pkg/modules/appointment"
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
