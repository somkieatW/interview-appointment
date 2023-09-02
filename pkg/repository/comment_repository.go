package repository

import (
	"context"
	"github.com/somkieatW/interview-appointment/pkg/core/utils"
	"github.com/somkieatW/interview-appointment/pkg/domain"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(ctx context.Context, obj *domain.Comments) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db}
}

func (r *commentRepository) Create(ctx context.Context, obj *domain.Comments) error {
	db := r.db.WithContext(ctx)
	err := db.Create(&obj).Error
	if err != nil {
		return utils.DbError(err)
	}
	return nil
}
