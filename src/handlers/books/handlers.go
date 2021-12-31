package books

import (
	"github.com/gen2brain/go-fitz"
	database "gitlab.com/kian00sh/rockbooks-be/src/database/postgresql"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/grapherrors"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/pagination"
)

// Books

func (book *Book) CreateBook() (*Book, error) {
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
	return book, nil
}

func (book *Book) GetBook() (*Book, error) {
	var bookResult *Book
	result := database.DB.Where("id = ?", book.ID).First(&bookResult)
	if result.Error != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در دریافت کتاب پیش آمده است!", result.Error)
	}
	return bookResult, nil
}

func (books *Book) GetBooks() ([]*Book, *pagination.Pagination, error) {
	var booksResult []*Book
	var paginationResult pagination.Pagination
	paginationResult.PaginationInput = books.PaginationInput
	result := database.DB.Scopes(pagination.Paginate(booksResult, &paginationResult, database.DB)).Find(&booksResult)
	if result.Error != nil {
		return nil, nil, grapherrors.ReturnGQLError("در دریافت کتاب ها مشکلی پیش آمده است", result.Error)
	}
	return booksResult, &paginationResult, nil
}

func (book *Book) GetBookAuthor() (*Author, error) {
	var authorResult *Author
	result := database.DB.Where("id = ?", book.AuthorID).First(&authorResult)
	if result.Error != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در دریافت نویسنده پیش آمده است!", result.Error)
	}
	return authorResult, nil
}

func (book *Book) GetBookPublisher() (*Publisher, error) {
	var publisherResult *Publisher
	result := database.DB.Where("id = ?", book.PublisherID).First(&publisherResult)
	if result.Error != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در دریافت ناشر پیش آمده است!", result.Error)
	}
	return publisherResult, nil
}

// BookPage
func (bookPage *BookPage) GetBookPages() ([]*BookPage, *pagination.Pagination, error) {
	var bookPagesResult []*BookPage
	var paginationResult pagination.Pagination
	paginationResult.PaginationInput = bookPage.PaginationInput
	result := database.DB.Scopes(pagination.Paginate(bookPagesResult, &paginationResult, database.DB)).Where("book_id = ?", bookPage.BookID).Find(&bookPagesResult)
	if result.Error != nil {
		return nil, nil, grapherrors.ReturnGQLError("در دریافت صفحات کتاب مشکلی پیش آمده!", result.Error)
	}
	return bookPagesResult, &paginationResult, nil
}

// Author

func (author *Author) CreateAuthor() (*Author, error) {
	result := database.DB.Create(&author)
	if result.Error != nil {
		return nil, grapherrors.ReturnGQLError("این نویسنده قبلا ثبت شده است", result.Error)
	}
	return author, nil
}

func (author *Author) GetAuthor() (*Author, error) {
	var authorResult *Author
	result := database.DB.Where("id = ?", author.ID).First(&authorResult)
	if result.Error != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در دریافت نویسنده پیش آمده است!", result.Error)
	}
	// bookResult.auth
	return authorResult, nil
}

// Publisher

func (publisher *Publisher) CreatePublisher() (*Publisher, error) {
	result := database.DB.Create(&publisher)
	if result.Error != nil {
		return nil, grapherrors.ReturnGQLError("این ناشر قبلا ثبت شده است", result.Error)
	}
	return publisher, nil
}

func (book *Publisher) GetPublisher() (*Publisher, error) {
	var publisherResult *Publisher
	result := database.DB.Where("id = ?", book.ID).First(&publisherResult)
	if result.Error != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در دریافت ناشر پیش آمده است!", result.Error)
	}
	// publisherResult.auth
	return publisherResult, nil
}

// audios
// pages
