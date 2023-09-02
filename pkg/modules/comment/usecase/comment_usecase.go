package usecase

import (
	"github.com/somkieatW/interview-appointment/pkg/core/registry"
	"github.com/somkieatW/interview-appointment/pkg/core/registry/core"
	"github.com/somkieatW/interview-appointment/pkg/modules/comment"
)

type commentUseCase struct {
	CoreRegistry       *core.CoreRegistry
	RepositoryRegistry *registry.RepositoryRegistry
}

func NewCommentUseCase(coreRegistry *core.CoreRegistry, repositoryRegistry *registry.RepositoryRegistry) comment.CommentUseCase {
	return &commentUseCase{
		CoreRegistry:       coreRegistry,
		RepositoryRegistry: repositoryRegistry,
	}
}
