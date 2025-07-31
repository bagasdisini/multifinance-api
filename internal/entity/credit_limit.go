package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type CreditLimit struct {
	ID             string         `gorm:"type:varchar(36);primaryKey"`
	CustomerID     string         `gorm:"type:varchar(36);index"`
	TenorInMonth   uint8          `gorm:"type:tinyint;index"`
	LimitAmount    float64        `gorm:"type:decimal(15,2);index"`
	RemainingLimit float64        `gorm:"type:decimal(15,2);index"`
	CreatedAt      time.Time      `gorm:"type:timestamp"`
	UpdatedAt      time.Time      `gorm:"type:timestamp"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (u *CreditLimit) BeforeCreate(*gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return
}

func (u *CreditLimit) TableName() string {
	return "credit_limits"
}
