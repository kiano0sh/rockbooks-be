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

type AudioBook struct {
	gorm.Model
	UserID       uint
	BookID       uint
	CursorStarts uint
	CursorEnds   uint
}

type Page struct {
	gorm.Model
	Content    string
	PageNumber int
	BookID     uint
}

type Book struct {
	gorm.Model
	Name        string `gorm:"size:128"`
	Pages       []Page
	AudioBooks  []AudioBook
	AuthorID    uint
	PublisherID uint
}
