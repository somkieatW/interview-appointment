package repository

import (
	"context"
	"github.com/somkieatW/interview-appointment/pkg/core/utils"
	"github.com/somkieatW/interview-appointment/pkg/domain"
	"github.com/somkieatW/interview-appointment/pkg/modules/appointment/models"
	"gorm.io/gorm"
)

type AppointmentRepository interface {
	IsExisted(ctx context.Context, id string) bool
	IsNotExisted(ctx context.Context, id string) bool
	List(ctx context.Context, obj *models.AppointmentListRequest) (*[]domain.Appointment, error)
	Info(ctx context.Context, obj *models.AppointmentInfoRequest) (*models.AppointmentInfoData, error)
	Update(ctx context.Context, obj *domain.Appointment) error
}

type appointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
	return &appointmentRepository{db}
}

func (r *appointmentRepository) IsExisted(ctx context.Context, id string) bool {
	appointment := &domain.Appointment{}
	db := r.db.WithContext(ctx)
	db = db.First(&domain.Appointment{}, id)

	if utils.IsEmpty(appointment.ID) {
		return false
	}

	return true
}

func (r *appointmentRepository) IsNotExisted(ctx context.Context, id string) bool {
	return !r.IsExisted(ctx, id)
}

func (r *appointmentRepository) List(ctx context.Context, obj *models.AppointmentListRequest) (*[]domain.Appointment, error) {
	appointment := &[]domain.Appointment{}

	db := r.db.WithContext(ctx)
	db = db.Model(appointment)

	db.Limit(obj.PageSize).Offset(obj.Offset)

	db = db.Preload("comments").Find(&appointment)
	if err := db.Error; err != nil {
		return nil, utils.DbError(err)
	}
	return appointment, nil
}

func (r *appointmentRepository) Info(ctx context.Context, obj *models.AppointmentInfoRequest) (*models.AppointmentInfoData, error) {

	appointment := &models.AppointmentInfoData{}
	db := r.db.WithContext(ctx)

	err := db.
		Joins("JOIN users ON appointment.user_id = users.id").
		Preload("comments").
		Select("appointment.*, users.display_name AS display_name").
		Where("appointment.id = ?", obj.ID).
		First(appointment).Error

	if err != nil {
		return nil, utils.DbError(err)
	}
	return appointment, nil
}

func (r *appointmentRepository) Update(ctx context.Context, obj *domain.Appointment) error {
	db := r.db.WithContext(ctx)

	err := db.Updates(obj).Error
	if err != nil {
		return utils.DbError(err)
	}

	return nil
}
