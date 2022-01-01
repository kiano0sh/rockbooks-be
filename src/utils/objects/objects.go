package objects

import (
	"gitlab.com/kian00sh/rockbooks-be/graph/model"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/pagination"
)

func PaginationInputToPaginationOutput(paginationInput model.PaginationInput) *pagination.PaginationOutput {
	return &pagination.PaginationOutput{Limit: *paginationInput.Limit, Page: *paginationInput.Page, SortBy: paginationInput.SortBy.String(), SortOrder: paginationInput.SortOrder.String()}
}
