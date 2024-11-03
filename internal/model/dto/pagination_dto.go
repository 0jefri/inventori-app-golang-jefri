package dto

// Ini buat paging di taruh di parameter
type PaginationParam struct {
	Page   int
	Offset int
	Limit  int
}

// ini buat paging di taruh di return
type PaginationQuery struct {
	Page int
	Take int
	Skip int
}

// ini buat di taruh di response
type Paging struct {
	Page        int `json:"page"`
	RowsPerPage int `json:"rowsPerPage"`
	TotalRows   int `json:"totalRows"`
	TotalPages  int `json:"totalPages"`
}

// example pagination -> product 100
// Paging {Page: 1, RowsPerPage: 10, TotalRows: 100, TotalPages: 10}
