package migrations

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	ID          uint64         `json:"id" gorm:"primaryKey"`
	Uuid        string         `json:"uuid" gorm:"size:50;uniqueIndex;not null"`
	Email       string         `json:"email" gorm:"size:255;not null;uniqueIndex"`
	Password    string         `json:"password"`
	PhoneNumber string         `json:"phoneNumber" gorm:"size:15;not null"`
	Role        int8           `json:"role" gorm:"not null"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
