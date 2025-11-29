package models

import (
	"time"
)

type Role struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Role      string `gorm:"type:varchar(50);uniqueIndex:idx_role" json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
