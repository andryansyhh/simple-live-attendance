package models

import (
	"time"
)

type Attendance struct {
	Uuid      string    `gorm:"primaryKey;uuid"`
	UserUuid  string    `gorm:"user_uuid"`
	ClockIn   time.Time `gorm:"clock_in"`
	ClockOut  time.Time `gorm:"clock_out"`
	Photo     string    `gorm:"photo"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

type AttendanceRequest struct {
	Uuid     string `json:"uuid"`
	UserUuid string `json:"user_uuid" binding:"required"`
	Photo    string `json:"photo" binding:"required"`
}
