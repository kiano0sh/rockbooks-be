package books

import (
	"github.com/gen2brain/go-fitz"
	"gitlab.com/kian00sh/rockbooks-be/graph/model"
	database "gitlab.com/kian00sh/rockbooks-be/src/database/postgresql"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/grapherrors"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/pagination"
)

// Books

func (book *Book) GetBook() (*model.Book, error) {
	var bookResult *Book
	result := database.DB.Where("id = ?", book.ID).First(&bookResult)
	if result.Error != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در دریافت کتاب پیش آمده است!", result.Error)
	}
	// bookResult.auth
	return &model.Book{ID: bookResult.ID, Name: bookResult.Name, CreatedAt: bookResult.CreatedAt.String()}, nil
}

func (books *Book) GetBooks() ([]*model.Book, *pagination.Pagination, error) {
	var booksResult []*Book
	var paginationResult pagination.Pagination
	paginationResult.PaginationInput = books.PaginationInput
	result := database.DB.Scopes(pagination.Paginate(booksResult, &paginationResult, database.DB)).Find(&booksResult)
	if result.Error != nil {
		return nil, nil, grapherrors.ReturnGQLError("در دریافت کتاب ها مشکلی پیش آمده است", result.Error)
	}
	var typedBooksResult []*model.Book
	for _, book := range booksResult {
		typedBooksResult = append(typedBooksResult, &model.Book{ID: book.ID, Name: book.Name, CreatedAt: book.CreatedAt.String()})
	}
	return typedBooksResult, &paginationResult, nil
}

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
	return &model.Book{ID: book.ID, Name: book.Name, CreatedAt: book.CreatedAt.String()}, nil
}

// Pages
func (bookPage *BookPage) GetBookPages() ([]*model.BookPage, *pagination.Pagination, error) {
	var bookPagesResult []*model.BookPage
	var paginationResult pagination.Pagination
	paginationResult.PaginationInput = bookPage.PaginationInput
	result := database.DB.Scopes(pagination.Paginate(bookPagesResult, &paginationResult, database.DB)).Where("book_id = ?", bookPage.BookID).Find(&bookPagesResult)
	if result.Error != nil {
		return nil, nil, grapherrors.ReturnGQLError("در دریافت صفحات کتاب مشکلی پیش آمده!", result.Error)
	}
	return bookPagesResult, &paginationResult, nil
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
