package repository

import (
	"context"
	"github.com/somkieatW/interview-appointment/pkg/core/constants"
	"github.com/somkieatW/interview-appointment/pkg/core/utils"
	"github.com/somkieatW/interview-appointment/pkg/domain"
	"github.com/somkieatW/interview-appointment/pkg/models"
	"gorm.io/gorm"
)

type AppointmentRepository interface {
	IsExisted(ctx context.Context, id string) bool
	IsNotExisted(ctx context.Context, id string) bool
	List(ctx context.Context, obj *models.AppointmentListRequest) (*[]models.AppointmentListData, error)
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

	db.First(appointment, "id = ?", id)

	if utils.IsEmpty(appointment.ID) {
		return false
	}

	return true
}

func (r *appointmentRepository) IsNotExisted(ctx context.Context, id string) bool {
	return !r.IsExisted(ctx, id)
}

func (r *appointmentRepository) List(ctx context.Context, obj *models.AppointmentListRequest) (*[]models.AppointmentListData, error) {
	appointment := &[]models.AppointmentListData{}

	db := r.db.WithContext(ctx)
	db = db.Model(&domain.Appointment{})
	db.Limit(obj.PageSize).Offset(obj.Offset)

	err := db.Joins("JOIN user ON appointment.created_by = user.id").
		Select("appointment.*, user.display_name").
		Where("appointment.status = ?", constants.StatusActive).
		Order("created_at asc").
		Find(&appointment).Error
	if err != nil {
		return nil, utils.DbError(err)
	}
	return appointment, nil
}

func (r *appointmentRepository) Info(ctx context.Context, obj *models.AppointmentInfoRequest) (*models.AppointmentInfoData, error) {

	appointment := &models.AppointmentInfoData{}
	db := r.db.WithContext(ctx)

	if err := db.Model(&domain.Appointment{}).
		Select("appointment.*, user.display_name as created_by").
		Joins("LEFT JOIN user ON user.id = appointment.created_by").
		Where("appointment.id = ?", obj.ID).
		First(&appointment).Error; err != nil {
		return nil, err
	}

	comments := &[]models.Comments{}
	if err := db.Model(&models.Comments{}).
		Select("comments.*, user.display_name as comment_by").
		Joins("LEFT JOIN user on user.id = comments.user_id").
		Where("appointment_id = ?", obj.ID).
		Order("created_at desc").
		Find(&comments).Error; err != nil {
		return nil, err
	}

	appointment.Comments = comments

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
