package books

import (
	"fmt"

	"github.com/gen2brain/go-fitz"
	"gitlab.com/kian00sh/rockbooks-be/graph/model"
	database "gitlab.com/kian00sh/rockbooks-be/src/database/postgresql"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/grapherrors"
)

func (book *Book) CreateBook() (*model.Book, error) {
	doc, err := fitz.NewFromReader(book.BookFile.File)
	if err != nil {
		panic(err)
	}
	defer doc.Close()

	// Extract pages as text
	for n := 0; n < doc.NumPage(); n++ {
		text, err := doc.Text(n)
		if err != nil {
			panic(err)
		}
		fmt.Printf(text)
	}
	return &model.Book{}, nil
}

func (author *Author) CreateAuthor() (*model.Author, error) {
	result := database.DB.Create(&author)
	if result.Error != nil {
		return nil, grapherrors.ReturnGQLError("این نویسنده قبلا ثبت شده است", result.Error)
	}
	var authorBooks []*model.Book
	return &model.Author{ID: author.ID, Name: author.Name, Books: authorBooks}, nil
}

func (publisher *Publisher) CreatePublisher() (*model.Publisher, error) {
	result := database.DB.Create(&publisher)
	if result.Error != nil {
		return nil, grapherrors.ReturnGQLError("این ناشر قبلا ثبت شده است", result.Error)
	}
	var publisherBooks []*model.Book
	return &model.Publisher{ID: publisher.ID, Name: publisher.Name, Books: publisherBooks}, nil
}
