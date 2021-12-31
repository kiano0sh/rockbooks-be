package books

import (
	"github.com/99designs/gqlgen/graphql"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/pagination"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID    int64
	Name  string `gorm:"size:128;uniqueIndex"`
	Books []Book
}

type Publisher struct {
	gorm.Model
	ID    int64
	Name  string `gorm:"size:128;uniqueIndex"`
	Books []Book
}

type BookAudio struct {
	gorm.Model
	ID           int64
	Audio        string `gorm:"size:256"`
	UserID       int64
	BookID       int64
	CursorStarts int64
	CursorEnds   int64
}

type BookPage struct {
	gorm.Model
	pagination.PaginationInput `gorm:"-"`
	ID                         int64
	Content                    string
	PageNumber                 int
	BookID                     int64
}

type Book struct {
	gorm.Model
	pagination.PaginationInput `gorm:"-"`
	ID                         int64
	Name                       string `gorm:"size:128"`
	Cover                      string `gorm:"size:256"`
	Wallpaper                  string `gorm:"size:256"`
	Pages                      []BookPage
	AudioBooks                 []BookAudio
	AuthorID                   int64
	PublisherID                int64
	BookFile                   graphql.Upload `gorm:"-"`
	WallpaperFile              graphql.Upload `gorm:"-"`
	CoverFile                  graphql.Upload `gorm:"-"`
}
