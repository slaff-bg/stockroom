package users

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Email     string    `gorm:"size:64;not null"`
	Passwd    string    `gorm:"size:64;not null"`
	FirstName string    `gorm:"size:32;not null"`
	LastName  string    `gorm:"size:32;not null"`
	CreatedAt time.Time `gorm:"type:date-time"`
	UpdatedAt time.Time `gorm:"type:date-time"`
}
