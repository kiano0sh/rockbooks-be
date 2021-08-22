package books

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name       string      `gorm:"size:128;uniqueIndex"`
	Publishers []Publisher `gorm:"many2many:author_publishers;"`
	Books      []Book
}

type Publisher struct {
	gorm.Model
	Name    string      `gorm:"size:128;uniqueIndex"`
	Authors []Publisher `gorm:"many2many:author_publishers;"`
	Books   []Book
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
}
