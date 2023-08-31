package repository

import (
	"context"
	"github.com/somkieatW/candidate/pkg/domain"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(ctx context.Context, obj *domain.Comment) (*domain.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db}
}

func (r *commentRepository) Create(ctx context.Context, obj *domain.Comment) (*domain.Comment, error) {
	comment := &domain.Comment{}
	db := r.db.WithContext(ctx)
	db = db.Create(&obj)

	if err := db.Error; err != nil {
		return nil, DbError(err)
	}
	return comment, nil
}
