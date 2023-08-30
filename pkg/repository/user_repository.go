package repository

import (
	"context"
	"github.com/somkieatW/candidate/pkg/domain"
	"github.com/somkieatW/candidate/pkg/modules/user/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Info(ctx context.Context, obj *models.UserInfoRequest) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Info(ctx context.Context, obj *models.UserInfoRequest) (*domain.User, error) {
	user := &domain.User{}
	db := r.db.WithContext(ctx)
	db = db.First(&obj)

	if err := db.Error; err != nil {
		return nil, DbError(err)
	}
	return user, nil
}
