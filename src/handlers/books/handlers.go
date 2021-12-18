package books

import (
	"fmt"

	"github.com/gen2brain/go-fitz"
	"gitlab.com/kian00sh/rockbooks-be/graph/model"
	database "gitlab.com/kian00sh/rockbooks-be/src/database/postgresql"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/grapherrors"
)

func (book *Book) CreateBook() {
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
		fmt.Println(text)
	}
}

func (author *Author) CreateAuthor() (*model.Author, error) {
	result := database.DB.Create(&author)
	if result.Error != nil {
		return nil, grapherrors.ReturnGQLError("این نویسنده قبلا ثبت شده است", result.Error)
	}
	var authorBooks []*model.Book
	authorBooksResult := database.DB.Where("author_id = ?", author.ID).Find(authorBooks)
	if authorBooksResult.Error != nil {
		return nil, grapherrors.ReturnGQLError("در هنگام دیافت کتاب ‌های این نویسنده مشکلی پیش آمد", result.Error)
	}
	return &model.Author{Name: author.Name, Books: authorBooks}, nil
}
