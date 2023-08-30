package repository

import (
	"context"
	"database/sql"
	"github.com/somkieatW/candidate/pkg/domain"
	"github.com/somkieatW/candidate/pkg/repository/repomodels"

	"gorm.io/gorm"
)

type AppointmentRepository interface {
	List(ctx context.Context, obj *repomodels.AppointmentListCriteria) (*[]domain.Appointment, error)
}

type appointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
	return &appointmentRepository{db}
}

func (r *appointmentRepository) List(ctx context.Context, obj *repomodels.AppointmentListCriteria) (*[]domain.Appointment, error) {
	appointment := &[]domain.Appointment{}

	db := r.db.WithContext(ctx)
	db = db.Model(appointment)

	db.Limit(obj.Limit.PageSize).Offset(obj.Limit.Offset)

	db = db.Find(&appointment)
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
