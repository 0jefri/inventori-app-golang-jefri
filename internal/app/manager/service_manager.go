package manager

import (
	"github.com/inventori-app-jeff/internal/app/service"
)

type ServiceManager interface {
	UserService() service.UserService
	AuthService() service.AuthService
	ProductService() service.ProductService
	TransactionService() service.TransactionService
	CategoryService() service.CategoryService
}

type serviceManager struct {
	repoManager RepoManager
}

func NewServiceManager(repo RepoManager) ServiceManager {
	return &serviceManager{
		repoManager: repo,
	}
}

func (m *serviceManager) UserService() service.UserService {
	return service.NewUserService(m.repoManager.UserRepo())
}

func (m *serviceManager) AuthService() service.AuthService {
	return service.NewAuthService(m.UserService())
}

func (m *serviceManager) ProductService() service.ProductService {
	return service.NewProductService(m.repoManager.ProductRepo())
}

func (m *serviceManager) TransactionService() service.TransactionService {
	return service.NewTransactionService(m.repoManager.TransactionRepo(), m.repoManager.ProductRepo())
}

func (m *serviceManager) CategoryService() service.CategoryService {
	return service.NewCategoryService(m.repoManager.CategoryRepo(), m.repoManager.ProductRepo())
}
