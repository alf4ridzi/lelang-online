package seeders

import (
	"lelang-online-api/models"

	"gorm.io/gorm"
)

var RoleSeed = []models.Role{
	{
		Role: "admin",
	},
	{
		Role: "petugas",
	},
	{
		Role: "user",
	},
}

func RoleSeeeder(db *gorm.DB) {
	for _, role := range RoleSeed {
		db.FirstOrCreate(&role, models.Role{Role: role.Role})
	}
}
