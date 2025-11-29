package migrations

import "gorm.io/gorm"

func migrate(db *gorm.DB, model []any) error {
	return db.AutoMigrate(model...)
}
