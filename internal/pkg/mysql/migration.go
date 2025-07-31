package mysql

import (
	"fmt"
	"github.com/bagasdisini/multifinance-api/internal/entity"
)

func RunMigration() {
	var err error
	err = DB.AutoMigrate(
		&entity.Customer{},
		&entity.CreditLimit{},
		&entity.Transaction{},
	)

	if err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}
}
