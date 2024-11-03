package service

import (
	"errors"

	"github.com/inventori-app-jeff/internal/app/repository"
	"github.com/inventori-app-jeff/internal/model"

	"github.com/inventori-app-jeff/utils/constants"
)

type TransactionService interface {
	ReceiveTransaction(payload *model.Transaction) (*model.Transaction, error)
	CreateSendTransaction(payload *model.Transaction) (*model.Transaction, error)
	ListTransactions() ([]*model.Transaction, error)
}

type transactionService struct {
	repo        repository.TransactionRepository
	productRepo repository.ProductRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository, productRepo repository.ProductRepository) TransactionService {
	return &transactionService{
		repo:        transactionRepo,
		productRepo: productRepo,
	}
}

func (s *transactionService) ReceiveTransaction(payload *model.Transaction) (*model.Transaction, error) {
	if payload.Amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}
	if payload.TransactionType != constants.ReceiveProductType {
		return nil, errors.New("invalid transaction type")
	}

	transaction, err := s.repo.CreateReceiveTransaction(payload)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *transactionService) CreateSendTransaction(payload *model.Transaction) (*model.Transaction, error) {
	// Validasi input transaksi
	if payload == nil {
		return nil, errors.New("transaction payload cannot be nil")
	}
	if payload.Amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}
	if payload.ProductID == "" {
		return nil, errors.New("product ID is required")
	}

	// Memanggil repository untuk membuat transaksi pengiriman
	transaction, err := s.repo.CreateSendTransaction(payload)
	if err != nil {
		return nil, err // Mengembalikan error jika ada kesalahan di repository
	}

	return transaction, nil
}

func (s *transactionService) ListTransactions() ([]*model.Transaction, error) {
	// Memanggil metode List pada repository untuk mengambil semua transaksi
	transactions, err := s.repo.List()
	if err != nil {
		return nil, err // Mengembalikan error jika terdapat kesalahan pada repository
	}

	return transactions, nil
}
