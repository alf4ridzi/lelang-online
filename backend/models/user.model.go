package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	RoleID    int    `json:"-"`
	Role      Role   `json:"role"`
	Name      string `json:"name"`
	Username  string `gorm:"type:varchar(255);uniqueIndex:idx_username" json:"username"`
	Password  string `json:",omitempty"`
	Phone     string `gorm:"type:varchar(255);uniqueIndex:idx_phone" json:"phone"`
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
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirmpassword" binding:"required,min=6"`
	Phone           string `json:"phone"`
}
