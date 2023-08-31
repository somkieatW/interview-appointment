package usecase

import (
	"github.com/somkieatW/candidate/pkg/core/registry"
	"github.com/somkieatW/candidate/pkg/core/registry/core"
	"github.com/somkieatW/candidate/pkg/modules/comment"
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
