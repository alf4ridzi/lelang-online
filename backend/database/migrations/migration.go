package migrations

import "gorm.io/gorm"

func Migrate(db *gorm.DB, model []any) error {
	return db.AutoMigrate(model...)
}
