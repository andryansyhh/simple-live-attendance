package usecase

import (
	"errors"
	"simple-live-attendance/src/models"
	"simple-live-attendance/src/repository"
	"simple-live-attendance/src/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AttendanceUsecase struct {
	db             *gorm.DB
	attendanceRepo *repository.AttendanceRepository
}

func NewUsecase(db *gorm.DB) *AttendanceUsecase {
	return &AttendanceUsecase{
		db:             db,
		attendanceRepo: repository.NewAttendanceRepository(db),
	}
}

func (u *AttendanceUsecase) ClockIn(request models.AttendanceRequest) error {
	// calculate similarity if < 80% return error
	similarity, err := utils.CalculateSimilarity()
	if similarity < 80 {
		err = errors.New("attendance submission not valid, please retake the photo")
		return err
	}
	if err != nil {
		return err
	}

	attendance := models.Attendance{
		Uuid:      uuid.New().String(),
		UserUuid:  request.UserUuid,
		ClockIn:   time.Now(),
		Photo:     request.Photo,
		CreatedAt: time.Now(),
	}

	if err := u.attendanceRepo.CreateAttendance(&attendance); err != nil {
		return errors.New("failed to create attendance")
	}

	return nil
}

func (u *AttendanceUsecase) ClockOut(request models.AttendanceRequest) error {
	// calculate similarity if < 80% return error
	similarity, err := utils.CalculateSimilarity()
	if similarity < 80 {
		err = errors.New("attendance submission not valid, please retake the photo")
		return err
	}
	if err != nil {
		return err
	}

	attendance := models.Attendance{
		Uuid:      request.Uuid,
		UserUuid:  request.UserUuid,
		ClockOut:  time.Now(),
		Photo:     request.Photo,
		UpdatedAt: time.Now(),
	}

	if err := u.attendanceRepo.UpdateAttendance(&attendance); err != nil {
		return errors.New("failed to create attendance")
	}

	return nil
}
