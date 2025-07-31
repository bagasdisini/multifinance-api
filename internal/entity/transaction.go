package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID           string         `gorm:"type:varchar(36);primaryKey"`
	CustomerID   string         `gorm:"type:varchar(36);index"`
	OTR          float64        `gorm:"type:decimal(15,2)"`
	AdminFee     float64        `gorm:"type:decimal(15,2)"`
	Interest     float64        `gorm:"type:decimal(15,2)"`
	TenorInMonth uint8          `gorm:"type:tinyint;index"`
	AssetName    string         `gorm:"type:varchar(255)"`
	Status       string         `gorm:"type:varchar(20)"`
	CreatedAt    time.Time      `gorm:"type:timestamp"`
	UpdatedAt    time.Time      `gorm:"type:timestamp"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (u *Transaction) BeforeCreate(*gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return
}

func (u *Transaction) TableName() string {
	return "transactions"
}
