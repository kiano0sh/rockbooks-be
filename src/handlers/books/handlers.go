package books

import (
	"fmt"

	"github.com/gen2brain/go-fitz"
	"gitlab.com/kian00sh/rockbooks-be/graph/model"
	database "gitlab.com/kian00sh/rockbooks-be/src/database/postgresql"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/grapherrors"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/pagination"
)

// Books

func (book *Book) CreateBook() (*model.Book, error) {
	theBook, err := fitz.NewFromReader(book.BookFile.File)
	if err != nil {
		grapherrors.ReturnGQLError("مشکلی در آغاز فرایند ثبت کتاب پیش آمده است، لطفا مجددا تلاش کنید", err)
	}
	defer theBook.Close()
	// Collect pages in an array of pages to be used for batch insert
	var pages []BookPage
	// Extract pages as text
	for pageNumber := 0; pageNumber < theBook.NumPage(); pageNumber++ {
		text, err := theBook.Text(pageNumber)
		if err != nil {
			// TODO delete the book
			grapherrors.ReturnGQLError("مشکلی در ثبت صفحات کتاب پیش آمده است، لطفا مجددا تلاش کنید", err)
		}
		pages = append(pages, BookPage{Content: text, PageNumber: pageNumber})
	}
	book.Pages = pages
	result := database.DB.Create(&book)
	if result.Error != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در ثبت کتاب پیش آمده است، لطفا مجددا تلاش کنید", result.Error)
	}
	// TODO Fetch the Author from its method and place it in the output
	// TODO Fetch the Publisher from its method and place it in the output
	return &model.Book{Name: book.Name}, nil
}

// Pages
func (bookPage *BookPage) GetPages() ([]*model.BookPage, error) {
	fmt.Printf("%d", bookPage.Limit)
	var pages []*model.BookPage
	database.DB.Scopes(pagination.Paginate(pages, &pagination.Pagination{PaginationInput: bookPage.PaginationInput}, database.DB)).Find(&pages)
	// cg.db.Scopes(paginate(categories,&pagination, cg.db)).Find(&categories)

	return nil, nil
}

// author
// publisher
// audios
// pages

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
