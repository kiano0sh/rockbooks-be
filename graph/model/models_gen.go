// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/99designs/gqlgen/graphql"
)

type Author struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Books []*Book `json:"books"`
}

type Book struct {
	Name      string     `json:"name"`
	Author    *Author    `json:"author"`
	Publisher *Publisher `json:"publisher"`
	CreatedAt string     `json:"createdAt"`
}

type BookAudio struct {
	CreatedBy    *User  `json:"createdBy"`
	Audio        string `json:"audio"`
	Book         *Book  `json:"book"`
	CursorStarts int    `json:"cursorStarts"`
	CursorEnds   int    `json:"cursorEnds"`
	CreatedAt    string `json:"createdAt"`
}

type BookPage struct {
	Content    string `json:"content"`
	PageNumber int    `json:"pageNumber"`
}

type CreateAuthorInput struct {
	Name string `json:"name"`
}

type CreateBookAudioInput struct {
	Audio        string `json:"audio"`
	BookID       int64  `json:"bookId"`
	CursorStarts int    `json:"cursorStarts"`
	CursorEnds   int    `json:"cursorEnds"`
}

type CreateBookInput struct {
	Name        string         `json:"name"`
	AuthorID    int64          `json:"authorId"`
	PublisherID int64          `json:"publisherId"`
	BookFile    graphql.Upload `json:"bookFile"`
}

type CreatePublisherInput struct {
	Name string `json:"name"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Pagination struct {
	Limit *int `json:"Limit"`
	Page  *int `json:"Page"`
}

type Publisher struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Books []*Book `json:"books"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type RegisterInput struct {
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type UpdateAuthorInput struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type UpdateBookAudioInput struct {
	ID           int64  `json:"id"`
	Audio        string `json:"audio"`
	BookID       int64  `json:"bookId"`
	CursorStarts int    `json:"cursorStarts"`
	CursorEnds   int    `json:"cursorEnds"`
}

type UpdateBookInput struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	AuthorID    int64  `json:"authorId"`
	PublisherID int64  `json:"publisherId"`
}

type UpdatePublisherInput struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type User struct {
	DisplayName string  `json:"displayName"`
	Email       string  `json:"email"`
	Avatar      *string `json:"avatar"`
}
