package common

import (
	"math"
	"strconv"

	"github.com/inventori-app-jeff/config"
	"github.com/inventori-app-jeff/internal/model/dto"
)

func GetPaginationParams(params dto.PaginationParam) dto.PaginationQuery {
	var (
		page, take, skip int
	)

	if params.Page > 0 {
		page = params.Page
	} else {
		page = 1
	}

	if params.Limit == 0 {
		n, _ := strconv.Atoi(config.Cfg.DefaultRowsPerPage)
		take = n
	} else {
		take = params.Limit
	}

	// rumus offset / rumus pagination
	// semisal default limit rows perpage nya 5
	// product => 10 | Page 1 => row 1 s.d 5
	// product => 10 | Page 2 => row 6 s.d 10
	// offset = (page - 1) * limit
	// offset = (1 - 1) * 5 ==== 0
	// offset = (2 - 1) * 5 ==== 5
	// SELECT * FROM product LIMIT 5 OFFSET 5

	if page > 0 {
		skip = (page - 1) * take
	} else {
		skip = 0
	}

	return dto.PaginationQuery{
		Page: page,
		Take: take,
		Skip: skip,
	}
}

// 21 / 5 === 4.xxx
// ceil (pembulatan keatas) e.g 4.2 == 4 | 4.6 == 5
func Paginate(page, limit, totalRows int) *dto.Paging {
	return &dto.Paging{
		Page:        page,
		RowsPerPage: limit,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(limit))),
	}
}

// urutan pembuatan pagination pada handson ini, terakhir perubahan di branch 03-with-db
// 1. create dto > create pagination_dto
// 3. .env > create DEFAULT_ROWS_PER_PAGE
// 2. utils > create pagination_util
// 3. base_repository > create interface BaseRepositoryPaging
// 4. implementasi kan dalam employee_repository file
// 5. implementasi employee_usecase
// 6. controller
