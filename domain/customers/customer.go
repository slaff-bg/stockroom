package clients

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	brand_name string    `gorm:"size:64;not null"`
	CreatedAt  time.Time `gorm:"type:date-time"`
	UpdatedAt  time.Time `gorm:"type:date-time"`
}
