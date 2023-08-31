package repository

import (
	"context"
	"database/sql"
	"github.com/somkieatW/candidate/pkg/core/utils"
	"github.com/somkieatW/candidate/pkg/domain"
	"github.com/somkieatW/candidate/pkg/modules/appointment/models"
	"gorm.io/gorm"
)

type AppointmentRepository interface {
	IsExisted(ctx context.Context, id string) bool
	IsNotExisted(ctx context.Context, id string) bool
	List(ctx context.Context, obj *models.AppointmentListRequest) (*[]domain.Appointment, error)
	Info(ctx context.Context, obj *models.AppointmentInfoRequest) (*domain.Appointment, error)
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

	db = db.Find(&appointment)
	if err := db.Error; err != nil {
		return nil, DbError(err)
	}
	return appointment, nil
}

func (r *appointmentRepository) Info(ctx context.Context, obj *models.AppointmentInfoRequest) (*domain.Appointment, error) {
	appointment := &domain.Appointment{}
	db := r.db.WithContext(ctx)
	db = db.Model(obj)
	db = db.First(&appointment)

	if err := db.Error; err != nil {
		return nil, DbError(err)
	}
	return appointment, nil
}

func DbError(err error) error {
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}
