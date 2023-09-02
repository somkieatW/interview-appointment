package repository

import (
	"context"
	"github.com/somkieatW/interview-appointment/pkg/core/utils"
	"github.com/somkieatW/interview-appointment/pkg/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	IsExisted(ctx context.Context, id string) bool
	IsNotExisted(ctx context.Context, id string) bool
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) IsExisted(ctx context.Context, id string) bool {
	user := &domain.User{}
	db := r.db.WithContext(ctx)

	db.First(user, "id = ?", id)

	if utils.IsEmpty(user.ID) {
		return false
	}

	return true
}

func (r *userRepository) IsNotExisted(ctx context.Context, id string) bool {
	return !r.IsExisted(ctx, id)
}
