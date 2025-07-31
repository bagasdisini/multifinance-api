package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	ID          string         `gorm:"type:varchar(36);primaryKey"`
	NIK         string         `gorm:"type:varchar(16);unique;index;not null"`
	FullName    string         `gorm:"type:varchar(255)"`
	LegalName   string         `gorm:"type:varchar(255)"`
	BirthPlace  string         `gorm:"type:varchar(255)"`
	BirthDate   time.Time      `gorm:"type:date"`
	Salary      float64        `gorm:"type:decimal(15,2)"`
	KTPPhoto    string         `gorm:"type:varchar(255)"`
	SelfiePhoto string         `gorm:"type:varchar(255)"`
	CreatedAt   time.Time      `gorm:"type:timestamp"`
	UpdatedAt   time.Time      `gorm:"type:timestamp"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	CreditLimits []CreditLimit `gorm:"foreignKey:CustomerID"`
	Transactions []Transaction `gorm:"foreignKey:CustomerID"`
}

func (u *Customer) BeforeCreate(*gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return
}

func (u *Customer) TableName() string {
	return "customers"
}
