package books

import (
	database "gitlab.com/kian00sh/rockbooks-be/src/database/postgresql"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/grapherrors"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/pagination"
)

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
