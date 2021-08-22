package migrations

import (
	"gitlab.com/kian00sh/rockbooks-be/src/handlers/books"
	"gitlab.com/kian00sh/rockbooks-be/src/handlers/users"
	"gorm.io/gorm"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func InitMigrations(db *gorm.DB) {
	handleError(db.AutoMigrate(&users.User{}))
	handleError(db.AutoMigrate(&books.Publisher{}))
	handleError(db.AutoMigrate(&books.Author{}))
	handleError(db.AutoMigrate(&books.BookAudio{}))
	handleError(db.AutoMigrate(&books.BookPage{}))
	handleError(db.AutoMigrate(&books.Book{}))
}
