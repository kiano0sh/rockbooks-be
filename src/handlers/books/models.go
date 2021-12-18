package books

import (
	"github.com/99designs/gqlgen/graphql"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name  string `gorm:"size:128;uniqueIndex"`
	Books []Book
}

type Publisher struct {
	gorm.Model
	Name  string `gorm:"size:128;uniqueIndex"`
	Books []Book
}

type BookAudio struct {
	gorm.Model
	Audio        string `gorm:"size:256"`
	UserID       uint
	BookID       uint
	CursorStarts uint
	CursorEnds   uint
}

type BookPage struct {
	gorm.Model
	Content    string
	PageNumber int
	BookID     uint
}

type Book struct {
	gorm.Model
	Name        string `gorm:"size:128"`
	Pages       []BookPage
	AudioBooks  []BookAudio
	AuthorID    uint
	PublisherID uint
	BookFile    graphql.Upload `gorm:"-"`
}
