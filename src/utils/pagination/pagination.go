package pagination

import (
	"math"

	"gorm.io/gorm"
)

type PaginationInput struct {
	Limit int
	Page  int
	Sort  string
}

type PaginationOutput struct {
	Limit     int
	Page      int
	Total     int
	SortBy    string
	SortOrder string
}

type Pagination struct {
	PaginationInput
	TotalRows  int64
	TotalPages int
	Rows       interface{}
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}

func Paginate(totalRows int64, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

func GenerateSortStatement(pagination *PaginationOutput) string {
	return pagination.SortBy + " " + pagination.SortOrder
}

func CreatePaginationInput(pagination *PaginationOutput) PaginationInput {
	sortStatement := GenerateSortStatement(pagination)
	return PaginationInput{Page: pagination.Page, Limit: pagination.Limit, Sort: sortStatement}
}

func CreatePaginationResult(paginationValues *Pagination) *PaginationOutput {
	return &PaginationOutput{Limit: paginationValues.Limit, Page: paginationValues.Page, Total: paginationValues.TotalPages}
}
