package repository

import (
	"github.com/inventori-app-jeff/internal/model"
	"github.com/inventori-app-jeff/internal/model/dto"
	"github.com/inventori-app-jeff/utils/common"
	"github.com/inventori-app-jeff/utils/constants"
	"github.com/inventori-app-jeff/utils/exception"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository interface {
	BaseRepository[model.Product]
	BaseRepositoryPaging[model.Product]
	PagingByName(requestPaging dto.PaginationParam, name string) ([]*model.Product, *dto.Paging, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(payload *model.Product) (*model.Product, error) {

	product := model.Product{
		ID:       payload.ID,
		Name:     payload.Name,
		Quantity: payload.Quantity,
		Price:    payload.Price,
	}

	if err := r.db.Create(&product).Error; err != nil {
		return nil, exception.ErrFailedCreate
	}

	return &product, nil
}

func (r *productRepository) Paging(requestPaging dto.PaginationParam, queries ...string) ([]*model.Product, *dto.Paging, error) {

	products := []*model.Product{}

	paginationQuery := common.GetPaginationParams(requestPaging)

	var totalRows int64

	if err := r.db.Limit(paginationQuery.Take).Offset(paginationQuery.Skip).Find(&products).Error; err != nil {
		return nil, nil, err
	}

	var count int = int(totalRows)

	return products, common.Paginate(paginationQuery.Take, paginationQuery.Page, count), nil
}

func (r *productRepository) PagingByName(requestPaging dto.PaginationParam, name string) ([]*model.Product, *dto.Paging, error) {
	products := []*model.Product{}
	paginationQuery := common.GetPaginationParams(requestPaging)

	var totalRows int64

	query := r.db.Model(&model.Product{})
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if err := query.Count(&totalRows).Error; err != nil {
		return nil, nil, err
	}

	if err := query.Limit(paginationQuery.Take).
		Offset(paginationQuery.Skip).
		Find(&products).Error; err != nil {
		return nil, nil, err
	}

	count := int(totalRows)
	paging := common.Paginate(paginationQuery.Take, paginationQuery.Page, count)

	return products, paging, nil
}

func (r *productRepository) List() ([]*model.Product, error) {

	var products []*model.Product

	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) Get(id string) (*model.Product, error) {
	var product model.Product

	if err := r.db.Where(constants.WHERE_BY_ID, id).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (c *productRepository) Update(id string, payload *model.Product) (*model.Product, error) {

	product := model.Product{}

	err := c.db.Model(&product).Where(constants.WHERE_BY_ID, id).Clauses(clause.Returning{}).Updates(model.Product{
		ID:       payload.ID,
		Name:     payload.Name,
		Quantity: payload.Quantity,
		Price:    payload.Price,
	}).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (c *productRepository) Delete(id string) (*model.Product, error) {
	product := model.Product{}

	if err := c.db.Clauses(clause.Returning{}).Where(constants.WHERE_BY_ID, id).Delete(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}
