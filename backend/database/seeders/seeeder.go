package seeders

import "gorm.io/gorm"

func RunSeeder(db *gorm.DB) {
	RoleSeeeder(db)
	UserSeeder(db)
}
