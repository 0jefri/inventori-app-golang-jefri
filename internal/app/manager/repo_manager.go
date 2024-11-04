package manager

import (
	"github.com/inventori-app-jeff/internal/app/repository"
)

type RepoManager interface {
	UserRepo() repository.UserRepository
	ProductRepo() repository.ProductRepository
	TransactionRepo() repository.TransactionRepository
	CategoryRepo() repository.CategoryRepository
}

type repoManager struct {
	infraManager InfraManager
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infraManager: infra,
	}
}

func (m *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(m.infraManager.Conn())
}

func (m *repoManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(m.infraManager.Conn())
}

func (m *repoManager) TransactionRepo() repository.TransactionRepository {
	return repository.NewTransactionRepository(m.infraManager.Conn())
}

func (m *repoManager) CategoryRepo() repository.CategoryRepository {
	return repository.NewCategoryRepository(m.infraManager.Conn())
}
