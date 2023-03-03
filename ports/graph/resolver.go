package graph

import (
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	GDB *gorm.DB
}

// func (u *model.User) FetchAllUsers() {
// }
