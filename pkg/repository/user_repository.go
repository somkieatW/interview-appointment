package repository

import (
	"context"
	"github.com/somkieatW/candidate/pkg/core/utils"
	"github.com/somkieatW/candidate/pkg/domain"
	"github.com/somkieatW/candidate/pkg/modules/user/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	IsExisted(ctx context.Context, id string) bool
	IsNotExisted(ctx context.Context, id string) bool
	Info(ctx context.Context, obj *models.UserInfoRequest) (*domain.User, error)
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
	db = db.First(&domain.User{}, id)

	if utils.IsEmpty(user.ID) {
		return false
	}

	return true
}

func (r *userRepository) IsNotExisted(ctx context.Context, id string) bool {
	return !r.IsExisted(ctx, id)
}

func (r *userRepository) Info(ctx context.Context, obj *models.UserInfoRequest) (*domain.User, error) {
	user := &domain.User{}
	db := r.db.WithContext(ctx)
	db = db.First(&obj)

	if err := db.Error; err != nil {
		return nil, utils.DbError(err)
	}
	return user, nil
}
