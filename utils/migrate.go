package utils

import (
	"opinia/entities"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&entities.User{}, &entities.Postingan{}, &entities.PostType{},
	)
}
