package objects

import (
	"gitlab.com/kian00sh/rockbooks-be/graph/model"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/pagination"
)

func PaginationInputToPaginationOutput(paginationInput model.PaginationInput) *pagination.PaginationOutput {
	sortBy := "Id"
	if paginationInput.SortBy != nil {
		sortBy = paginationInput.SortBy.String()
	}
	sortOrder := "ASC"
	if paginationInput.SortOrder != nil {
		sortOrder = paginationInput.SortOrder.String()
	}
	return &pagination.PaginationOutput{Limit: *paginationInput.Limit, Page: *paginationInput.Page, SortBy: sortBy, SortOrder: sortOrder}
}
