package models

import (
	"time"
)

type Role struct {
	ID        uint   `gorm:"primaryKey"`
	Role      string `gorm:"uniqueIndex:idx_role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
