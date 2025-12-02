package seeders

import (
	"lelang-online-api/models"

	"gorm.io/gorm"
)

var ItemSeed = []models.Item{
	{
		Name:        "JAM TANGAN",
		Description: "Jam tangan keren",
		UserID:      1,
	},
}

func ItemSeeder(db *gorm.DB) {

}
