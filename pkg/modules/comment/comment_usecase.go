package comment

import (
	"context"
	"github.com/somkieatW/candidate/pkg/modules/comment/models"
)

type CommentUseCase interface {
	Create(ctx context.Context, request *models.CommentCreateRequest) (*models.CommentCreateResponse, error)
}
