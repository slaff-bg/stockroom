package customers

import (
	"time"

	"gorm.io/gorm"
)

// import (
// 	"time"

// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

//	type Customer struct {
//		gorm.Model
//		ID         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
//		brand_name string    `gorm:"size:64;not null"`
//		CreatedAt  time.Time `gorm:"type:date-time"`
//		UpdatedAt  time.Time `gorm:"type:date-time"`
//	}

type Customer struct {
	gorm.Model
	ID        string    `json:"id"`
	BrandName string    `json:"brand_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
