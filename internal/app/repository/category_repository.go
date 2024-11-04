package repository

import (
	"github.com/inventori-app-jeff/internal/model"
	"github.com/inventori-app-jeff/internal/model/dto"
	"github.com/inventori-app-jeff/utils/common"
	"github.com/inventori-app-jeff/utils/constants"
	"github.com/inventori-app-jeff/utils/exception"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(payload *model.Category) (*model.Category, error)
	// List() ([]*model.Category, error)
	ListCategory(id string) ([]*model.Category, error)
	BaseRepositoryPaging[model.Category]
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(payload *model.Category) (*model.Category, error) {
	category := model.Category{
		ID:           payload.ID,
		ProductID:    payload.ProductID,
		CategoryName: payload.CategoryName,
	}

	if err := r.db.Create(&category).Error; err != nil {
		return nil, exception.ErrFailedCreate
	}

	return &category, nil
}

func (r *categoryRepository) ListCategory(id string) ([]*model.Category, error) {
	categorys := []*model.Category{}

	if err := r.db.Where(constants.WHERE_BY_PRODUCT_ID, id).Find(&categorys).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return categorys, nil
}

// func (r *categoryRepository) List() ([]*model.Category, error) {

// 	var categorys []*model.Category

// 	if err := r.db.Find(&categorys).Error; err != nil {
// 		return nil, err
// 	}

// 	return categorys, nil
// }

func (r *categoryRepository) Paging(requestPaging dto.PaginationParam, queries ...string) ([]*model.Category, *dto.Paging, error) {

	categorys := []*model.Category{}

	paginationQuery := common.GetPaginationParams(requestPaging)

	var totalRows int64

	if err := r.db.Limit(paginationQuery.Take).Offset(paginationQuery.Skip).Find(&categorys).Error; err != nil {
		return nil, nil, err
	}

	var count int = int(totalRows)

	return categorys, common.Paginate(paginationQuery.Take, paginationQuery.Page, count), nil
}
