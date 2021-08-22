package migrations

import (
	"gitlab.com/kian00sh/rockbooks-be/src/handlers/users"
	"gorm.io/gorm"
)

func InitMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&users.User{})
	if err != nil {
		panic(err)
	}
}
