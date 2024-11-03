package manager

import (
	"github.com/inventori-app-jeff/internal/app/service"
)

type ServiceManager interface {
	UserService() service.UserService
	AuthService() service.AuthService
	// TransactionService() service.TransactionService
	// BillService() service.BillService
	// ContactService() service.ContactService
	// UserPictureService() service.UserPictureService
	// CardService() service.CardService
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

// func (m *serviceManager) TransactionService() service.TransactionService {
// 	return service.NewTransactionService(m.repoManager.TransactionRepo(), m.repoManager.UserRepo(), m.repoManager.CardRepo())
// }

// func (m *serviceManager) BillService() service.BillService {
// 	return service.NewBillService(m.repoManager.BillRepo(), m.repoManager.UserRepo(), m.repoManager.BillDetailsRepo())
// }

// func (m *serviceManager) ContactService() service.ContactService {
// 	return service.NewContactService(m.repoManager.ContactRepo(), m.repoManager.UserRepo())
// }

// func (m *serviceManager) UserPictureService() service.UserPictureService {
// 	return service.NewUserPictureService(m.repoManager.UserPictureRepo(), m.repoManager.FileRepo(), m.UserService())
// }

// func (m *serviceManager) CardService() service.CardService {
// 	return service.NewCardService(m.repoManager.CardRepo(), m.repoManager.UserRepo())
// }
