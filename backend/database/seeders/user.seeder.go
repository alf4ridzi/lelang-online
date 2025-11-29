package seeders

import (
	"lelang-online-api/models"
	"lelang-online-api/utils"
	"log"

	"gorm.io/gorm"
)

var UserSeed = []models.User{
	{
		RoleID:   1,
		Name:     "Alfaridzi",
		Username: "admin",
		Password: "admin123",
		Phone:    "0852342323",
	},
	{
		RoleID:   2,
		Name:     "Alfaridzi User",
		Username: "user",
		Password: "user123",
		Phone:    "0852342343",
	},
}

func UserSeeder(db *gorm.DB) {
	for _, user := range UserSeed {
		hashed, err := utils.GenerateHashBcrypt(user.Password)
		user.Password = hashed

		if err != nil {
			log.Fatal(err)
		}

		db.FirstOrCreate(&user, models.User{Username: user.Username})
	}
}
