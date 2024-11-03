package service

import (
	"github.com/inventori-app-jeff/internal/app/repository"
	"github.com/inventori-app-jeff/internal/model"
	"github.com/inventori-app-jeff/internal/model/dto"
	"github.com/inventori-app-jeff/utils/exception"
	"gorm.io/gorm"
)

type ProductService interface {
	RegisterNewProduct(payload *model.Product) (*dto.ProductResponse, error)
	FindProductByID(id string) (*dto.ProductResponse, error)
	FindAllProduct(requestPaging dto.PaginationParam, queries ...string) ([]*dto.ProductResponse, *dto.Paging, error)
	FindProductsByName(requestPaging dto.PaginationParam, name string) ([]*dto.ProductResponse, *dto.Paging, error)
	UpdateProductByID(id string, payload *model.Product) (*dto.ProductResponse, error)
	RemoveProduct(id string) (*dto.ProductResponse, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) RegisterNewProduct(payload *model.Product) (*dto.ProductResponse, error) {

	products, err := s.repo.List()

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	for _, product := range products {
		if product.Name == payload.Name {
			return nil, exception.ErrProductNameAlreadyExist
		}
	}

	product, err := s.repo.Create(payload)

	productResponse := dto.ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Quantity: product.Quantity,
		Price:    payload.Price,
	}

	return &productResponse, err
}

func (s *productService) FindProductByID(id string) (*dto.ProductResponse, error) {

	product, err := s.repo.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	productResponse := dto.ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Quantity: product.Quantity,
		Price:    product.Price,
	}

	return &productResponse, err
}

func (s *productService) FindAllProduct(requestPaging dto.PaginationParam, queries ...string) ([]*dto.ProductResponse, *dto.Paging, error) {

	products, paging, err := s.repo.Paging(requestPaging, queries...)

	if err != nil {
		return nil, nil, gorm.ErrRecordNotFound
	}

	var productResponses []*dto.ProductResponse

	for _, product := range products {

		productResponse := dto.ProductResponse{
			ID:       product.ID,
			Name:     product.Name,
			Quantity: product.Quantity,
			Price:    product.Price,
		}

		productResponses = append(productResponses, &productResponse)
	}

	return productResponses, paging, err
}

func (s *productService) FindProductsByName(requestPaging dto.PaginationParam, name string) ([]*dto.ProductResponse, *dto.Paging, error) {
	// Panggil PagingByName dari repository
	products, paging, err := s.repo.PagingByName(requestPaging, name)
	if err != nil {
		return nil, nil, err
	}

	// Konversi hasil query ke dalam bentuk DTO ProductResponse
	var productResponses []*dto.ProductResponse
	for _, product := range products {
		productResponse := &dto.ProductResponse{
			ID:       product.ID,
			Name:     product.Name,
			Quantity: product.Quantity,
			Price:    product.Price,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses, paging, nil
}

func (s *productService) RemoveProduct(id string) (*dto.ProductResponse, error) {

	product, err := s.repo.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	product, err = s.repo.Delete(product.ID)

	if err != nil {
		return nil, exception.ErrFailedDelete
	}

	productResponse := dto.ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Quantity: product.Quantity,
		Price:    product.Price,
	}

	return &productResponse, err
}

func (s *productService) UpdateProductByID(id string, payload *model.Product) (*dto.ProductResponse, error) {

	product, err := s.repo.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	product, err = s.repo.Update(product.ID, payload)

	if err != nil {
		return nil, exception.ErrFailedUpdate
	}

	productResponse := dto.ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Quantity: product.Quantity,
		Price:    product.Price,
	}

	return &productResponse, err
}
