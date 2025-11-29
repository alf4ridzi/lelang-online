package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	RoleID    int
	Role      Role
	Name      string
	Username  string `gorm:"uniqueIndex:idx_username"`
	Password  string
	Phone     string `gorm:"uniqueIndex:idx_phone"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=4"`
}

type Register struct {
	Username        string `json:"username" binding:"required"`
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirmpassword" binding:"required,min=6"`
}
