package mysql

import (
	"fmt"
	"github.com/bagasdisini/multifinance-api/internal/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DatabaseUsername,
		config.DatabasePassword,
		config.DatabaseHost,
		config.DatabasePort,
		config.DatabaseName,
	)

	var gormConfig = &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}}

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	sqlDB, err := DB.DB()
	if err != nil {
		panic(fmt.Sprintf("failed to get raw DB object: %v", err))
	}

	sqlDB.SetConnMaxIdleTime(30 * time.Minute)
	sqlDB.SetConnMaxLifetime(2 * time.Hour)
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(300)

	if err := sqlDB.Ping(); err != nil {
		panic(fmt.Sprintf("failed to ping database: %v", err))
	}
}
