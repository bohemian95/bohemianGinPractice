package database

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ProductId uint `gorm:"primaryKey;autoIncrement"`
	Name      string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
