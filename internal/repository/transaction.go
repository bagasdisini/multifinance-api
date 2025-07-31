package repository

import (
	"github.com/bagasdisini/multifinance-api/internal/entity"
	"gorm.io/gorm"
)

type Transaction interface {
	FindByID(id string) (entity.Transaction, error)
	FindAllByCustomerID(customerID string) ([]entity.Transaction, error)
	Create(transaction entity.Transaction) error
}

type TransactionRepository struct {
	Repository[entity.Transaction]
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		Repository[entity.Transaction]{
			DB: db,
		},
	}
}

func (r *TransactionRepository) FindByID(id string) (entity.Transaction, error) {
	var transaction entity.Transaction
	err := r.DB.Where("id = ?", id).First(&transaction).Error
	return transaction, err
}

func (r *TransactionRepository) FindAllByCustomerID(customerID string) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	err := r.DB.Where("customer_id = ?", customerID).Find(&transactions).Error
	return transactions, err
}

func (r *TransactionRepository) Create(transaction entity.Transaction) error {
	return r.DB.Create(&transaction).Error
}
