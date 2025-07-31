package repository

import (
	"github.com/bagasdisini/multifinance-api/internal/entity"
	"gorm.io/gorm"
)

type Customer interface {
	FindAll() ([]entity.Customer, error)
	FindByID(id string) (*entity.Customer, error)
	FindByEmail(email string) (*entity.Customer, error)
	Create(user *entity.Customer) error
}

type CustomerRepository struct {
	Repository[entity.Customer]
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		Repository[entity.Customer]{
			DB: db,
		},
	}
}

func (r *CustomerRepository) FindAll() ([]entity.Customer, error) {
	var users []entity.Customer
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *CustomerRepository) FindByID(id string) (*entity.Customer, error) {
	var user *entity.Customer
	err := r.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *CustomerRepository) FindByEmail(email string) (*entity.Customer, error) {
	var user *entity.Customer
	err := r.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *CustomerRepository) Create(user *entity.Customer) error {
	return r.DB.Create(user).Error
}
