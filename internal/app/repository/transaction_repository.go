// package repository

// import (
// 	"go/constant"

// 	"github.com/inventori-app-jeff/internal/model"
// 	"github.com/inventori-app-jeff/utils/constants"
// 	"gorm.io/gorm"
// )

// type TransactionRepository interface{
// 	CreateDepositTransaction(payload *model.Transaction) (*model.Transaction, error)
// }

// type transactionRepository struct {
// 	db *gorm.DB
// }

// func NewTransactionRepository(db *gorm.DB) TransactionRepository {
// 	return &transactionRepository{db: db}
// }

// func (r *transactionRepository) CreateDepositTransaction(payload *model.Transaction) (*model.Transaction, error) {
// 	transaction := model.Transaction{
// 		ID:              payload.ID,
// 		ProductID:          payload.ProductID,
// 		TransactionType: payload.TransactionType,
// 		Amount:          payload.Amount,
// 		Description:     payload.Description,
// 		Timestamp:       payload.Timestamp,
// 	}

// 	r.db.Transaction(func(tx *gorm.DB) error{
// 		product := model.Product{}

// 		if transaction.TransactionType == "receiveProduct" {
// 			if err := tx.Model(&product).Where(constant.WHERE_BY_PRODUCT_ID, transaction.ProductID).Select("quantity").First(&product).Error; err != nil {
// 				return gorm.ErrInvalidTransaction
// 			}
// 			product.Quantity += int(transaction.Amount)

// 			if err := tx.Model(&product).Where(constants.WHERE_BY_PRODUCT_ID, transaction.ProductID).Select("quantity").Updates(&product).Error; err != nil {
// 				return gorm.ErrInvalidTransaction
// 			}

// 			if err := tx.Create(&transaction).Error; err != nil {
// 				return gorm.ErrInvalidTransaction
// 			}
// 		} else {
// 			return gorm.ErrInvalidTransaction
// 		}
// 		return nil
// 	})
// 	return &transaction, nil
// }

package repository

import (
	"errors"

	"github.com/inventori-app-jeff/internal/model"
	"github.com/inventori-app-jeff/utils/constants"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateReceiveTransaction(payload *model.Transaction) (*model.Transaction, error)
	CreateSendTransaction(payload *model.Transaction) (*model.Transaction, error)
	FindByID(id string) (*model.Product, error)
	List() ([]*model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

// FindByID implements TransactionRepository.
func (r *transactionRepository) FindByID(id string) (*model.Product, error) {
	var product model.Product
	if err := r.db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *transactionRepository) CreateReceiveTransaction(payload *model.Transaction) (*model.Transaction, error) {
	transaction := model.Transaction{
		ID:              payload.ID,
		ProductID:       payload.ProductID,
		TransactionType: payload.TransactionType,
		Amount:          payload.Amount,
		Description:     payload.Description,
		Timestamp:       payload.Timestamp,
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		product := model.Product{}

		if transaction.TransactionType == "receiveProduct" {
			if err := tx.Model(&product).
				Where(constants.WHERE_BY_PRODUCT_ID, transaction.ProductID).
				Select("quantity").
				First(&product).Error; err != nil {
				return err
			}

			product.Quantity += int(transaction.Amount)
			if err := tx.Model(&product).
				Where(constants.WHERE_BY_PRODUCT_ID, transaction.ProductID).
				UpdateColumn("quantity", product.Quantity).Error; err != nil {
				return err
			}

			if err := tx.Create(&transaction).Error; err != nil {
				return err
			}
		} else {
			return gorm.ErrInvalidTransaction
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *transactionRepository) CreateSendTransaction(payload *model.Transaction) (*model.Transaction, error) {
	transaction := model.Transaction{
		ID:              payload.ID,
		ProductID:       payload.ProductID,
		TransactionType: payload.TransactionType,
		Amount:          payload.Amount,
		Description:     payload.Description,
		Timestamp:       payload.Timestamp,
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		product := model.Product{}

		// Validasi tipe transaksi
		if transaction.TransactionType == constants.SendProductType { // Pastikan menggunakan constant yang benar
			// Ambil data produk berdasarkan ProductID
			if err := tx.Model(&product).Where("id = ?", transaction.ProductID).Select("quantity").First(&product).Error; err != nil {
				return err
			}

			// Kurangi quantity produk
			if product.Quantity < int(transaction.Amount) {
				return errors.New("insufficient product quantity") // Menangani jika jumlah produk tidak mencukupi
			}
			product.Quantity -= int(transaction.Amount)

			// Update quantity produk
			if err := tx.Model(&product).Where("id = ?", transaction.ProductID).Update("quantity", product.Quantity).Error; err != nil {
				return err
			}

			// Simpan transaksi ke database
			if err := tx.Create(&transaction).Error; err != nil {
				return err
			}
		} else {
			return errors.New("invalid transaction type") // Error untuk tipe transaksi yang tidak valid
		}
		return nil
	})

	if err != nil {
		return nil, err // Return error jika transaksi gagal
	}

	return &transaction, nil
}

func (r *transactionRepository) List() ([]*model.Transaction, error) {
	transactions := []*model.Transaction{}

	if err := r.db.Find(&transactions).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return transactions, nil
}
