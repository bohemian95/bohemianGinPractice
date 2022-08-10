package database

import (
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	User      User    `gorm:"embedded"`
	Product   Product `gorm:"embedded"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
