package repository

import (
	"github.com/bagasdisini/multifinance-api/internal/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CreditLimit interface {
	FindAllByCustomerID(customerID string) ([]entity.CreditLimit, error)
	FindByCustomerIDAndTenorInMonth(customerID string, tenorInMonth uint8) (entity.CreditLimit, error)
	UpdateCreditLimit(customerID string, percentage float64) error
}

type CreditLimitRepository struct {
	Repository[entity.CreditLimit]
}

func NewCreditLimitRepository(db *gorm.DB) *CreditLimitRepository {
	return &CreditLimitRepository{
		Repository[entity.CreditLimit]{
			DB: db,
		},
	}
}

func (r *CreditLimitRepository) FindAllByCustomerID(customerID string) ([]entity.CreditLimit, error) {
	var creditLimits []entity.CreditLimit
	err := r.DB.Where("customer_id = ?", customerID).Find(&creditLimits).Error
	return creditLimits, err
}

func (r *CreditLimitRepository) FindByCustomerIDAndTenorInMonth(customerID string, tenorInMonth uint8) (entity.CreditLimit, error) {
	var creditLimit entity.CreditLimit
	err := r.DB.Where("customer_id = ? AND tenor_in_month = ?", customerID, tenorInMonth).First(&creditLimit).Error
	return creditLimit, err
}

func (r *CreditLimitRepository) FindByCustomerIDAndTenorInMonthWithLock(customerID string, tenorInMonth uint8) (entity.CreditLimit, error) {
	var cl entity.CreditLimit
	err := r.DB.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("customer_id = ? AND tenor_in_month = ?", customerID, tenorInMonth).
		First(&cl).Error
	return cl, err
}

func (r *CreditLimitRepository) UpdateCreditLimit(customerID string, percentage float64) error {
	return r.DB.
		Model(&entity.CreditLimit{}).
		Where("customer_id = ?", customerID).
		Update("remaining_limit", gorm.Expr("remaining_limit * ?", 1-percentage)).
		Error
}
