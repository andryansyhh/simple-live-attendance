package repository

import (
	"simple-live-attendance/src/models"

	"gorm.io/gorm"
)

type AttendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) *AttendanceRepository {
	return &AttendanceRepository{db: db}
}

func (repo *AttendanceRepository) CreateAttendance(attendance *models.Attendance) error {
	return repo.db.Create(attendance).Error
}

func (repo *AttendanceRepository) UpdateAttendance(attendance *models.Attendance) error {
	return repo.db.Model(&attendance).Updates(map[string]interface{}{
		"updated_at": attendance.UpdatedAt,
		"clock_out":  attendance.ClockOut,
		"photo":      attendance.Photo,
	}).Error
}
