package comment

import (
	"context"
	"github.com/somkieatW/interview-appointment/pkg/modules/comment/models"
)

type CommentUseCase interface {
	Create(ctx context.Context, request *models.CommentCreateRequest) (*models.CommentCreateResponse, error)
}
