package database

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserId    uint `gorm:"primaryKey;autoIncrement"`
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
