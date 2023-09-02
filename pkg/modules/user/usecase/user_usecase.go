package usecase

import (
	"github.com/somkieatW/interview-appointment/pkg/core/registry"
	"github.com/somkieatW/interview-appointment/pkg/core/registry/core"
	"github.com/somkieatW/interview-appointment/pkg/modules/user"
)

type userUseCase struct {
	CoreRegistry       *core.CoreRegistry
	RepositoryRegistry *registry.RepositoryRegistry
}

func NewUserUseCase(coreRegistry *core.CoreRegistry, repositoryRegistry *registry.RepositoryRegistry) user.UserUseCase {
	return &userUseCase{
		CoreRegistry:       coreRegistry,
		RepositoryRegistry: repositoryRegistry,
	}
}
