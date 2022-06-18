package fakers

import (
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"go_harians/database/migrations"
	"gorm.io/gorm"
	"time"
)

func UserFaker(db *gorm.DB) *migrations.Users {
	return &migrations.Users{
		Uuid:        uuid.New().String(),
		Email:       faker.Email(),
		Password:    "$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro0llC/.og/at2.uheWG/igi", //password
		PhoneNumber: faker.Phonenumber(),
		Role:        1,
		DeletedAt:   gorm.DeletedAt{},
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}
