package models

import (
	"time"
)

type Role struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Role      string `gorm:"uniqueIndex:idx_role" json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
