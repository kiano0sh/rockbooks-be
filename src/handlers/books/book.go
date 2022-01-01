package books

import (
	"io"
	"os"

	"github.com/gen2brain/go-fitz"
	database "gitlab.com/kian00sh/rockbooks-be/src/database/postgresql"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/consts"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/grapherrors"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/pagination"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/strings"
)

// Books

func (book *Book) CreateBook() (*Book, error) {
	// Handle Pages
	theBook, err := fitz.NewFromReader(book.BookFile.File)
	if err != nil {
		grapherrors.ReturnGQLError("مشکلی در آغاز فرایند ثبت کتاب پیش آمده است", err)
	}
	defer theBook.Close()
	// Collect pages in an array of pages to be used for batch insert
	var pages []BookPage
	// Extract pages as text
	for pageNumber := 0; pageNumber < theBook.NumPage(); pageNumber++ {
		text, err := theBook.Text(pageNumber)
		if err != nil {
			grapherrors.ReturnGQLError("مشکلی در ثبت صفحات کتاب پیش آمده است", err)
		}
		pages = append(pages, BookPage{Content: text, PageNumber: pageNumber + 1})
	}
	// Add pages to book object
	book.Pages = pages

	mainFilePath := consts.IMAGES_PATH + strings.NormalizeMediaName(book.Name)
	// Handle Book Cover
	coverPath, err := strings.ConcatExtensionToEnd(mainFilePath+"-cover", book.CoverFile.ContentType)
	if err != nil {
		return nil, err
	}
	touchedCoverFile, err := os.OpenFile(coverPath, consts.CREATE_FILE_FLAG, consts.CREATE_FILE_PERMISSION)
	if err != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در ثبت کاور کتاب پیش آمده است", err)
	}
	defer touchedCoverFile.Close()
	bytes, err := io.Copy(touchedCoverFile, book.CoverFile.File)
	_ = bytes

	// If error is not nil then panics
	if err != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در ثبت کاور کتاب پیش آمده است", err)
	}

	book.Cover = coverPath

	// Handle Book Wallpaper
	wallpaperPath, err := strings.ConcatExtensionToEnd(mainFilePath+"-wallpaper", book.CoverFile.ContentType)
	if err != nil {
		return nil, err
	}
	touchedWallpaperFile, err := os.OpenFile(wallpaperPath, consts.CREATE_FILE_FLAG, consts.CREATE_FILE_PERMISSION)
	if err != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در ثبت والپیپر کتاب پیش آمده است", err)
	}
	defer touchedWallpaperFile.Close()
	io.Copy(touchedWallpaperFile, book.WallpaperFile.File)
	book.Wallpaper = wallpaperPath

	// Create the book
	result := database.DB.Create(&book)
	if result.Error != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در ثبت کتاب پیش آمده است", result.Error)
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
