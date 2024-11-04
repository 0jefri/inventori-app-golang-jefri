package service

import (
	"github.com/inventori-app-jeff/internal/app/repository"
	"github.com/inventori-app-jeff/internal/model"
	"github.com/inventori-app-jeff/internal/model/dto"
	"github.com/inventori-app-jeff/utils/exception"
	"gorm.io/gorm"
)

type CategoryService interface {
	RegisterNewCategory(payload *model.Category) (*dto.CategoryResponse, error)
	FindAllCategory(id string) ([]*dto.CategoryResponse, error)
}

type categoryService struct {
	repo     repository.CategoryRepository
	prodRepo repository.ProductRepository
}

func NewCategoryService(repo repository.CategoryRepository, prodRepo repository.ProductRepository) CategoryService {
	return &categoryService{
		repo:     repo,
		prodRepo: prodRepo,
	}
}

func (s *categoryService) RegisterNewCategory(payload *model.Category) (*dto.CategoryResponse, error) {

	product, err := s.prodRepo.Get(payload.ProductID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	category, err := s.repo.Create(payload)

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	categoryResponses := dto.CategoryResponse{
		ID: category.ID,
		Product: model.Product{
			ID:       product.ID,
			Name:     product.Name,
			Quantity: product.Quantity,
			Price:    product.Price,
		},
		CategoryName: category.CategoryName,
	}

	return &categoryResponses, err

}

func (s *categoryService) FindAllCategory(id string) ([]*dto.CategoryResponse, error) {
	product, err := s.prodRepo.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	Categorys, err := s.repo.ListCategory(product.ID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	categoryResponses := []*dto.CategoryResponse{}

	for _, category := range Categorys {
		if category.ProductID == product.ID {
			categoryResponses = append(categoryResponses, &dto.CategoryResponse{
				ID: category.ID,
				Product: model.Product{
					ID:       product.ID,
					Name:     product.Name,
					Quantity: product.Quantity,
					Price:    product.Price,
				},
				CategoryName: category.CategoryName,
			})
		}
	}

	return categoryResponses, err
}
