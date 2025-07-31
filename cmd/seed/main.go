package seed

import (
	"fmt"
	"github.com/bagasdisini/multifinance-api/internal/entity"
	"github.com/bagasdisini/multifinance-api/internal/pkg/mysql"
	"github.com/google/uuid"
	"github.com/labstack/gommon/random"
	"time"
)

func RunSeed() {
	mysql.RunMigration()

	for _, customer := range customers {
		if err := mysql.DB.Create(&customer).Error; err != nil {
			return
		}
	}

	for _, creditLimit := range creditLimits {
		if err := mysql.DB.Create(&creditLimit).Error; err != nil {
			return
		}
	}

	fmt.Println("Seeding completed")
}

var customers = []entity.Customer{
	{
		ID:          uuid.New().String(),
		NIK:         random.String(16, random.Numeric),
		FullName:    "Budi",
		LegalName:   "Budi",
		BirthPlace:  "Jakarta",
		BirthDate:   time.Now(),
		Salary:      10_000_000,
		KTPPhoto:    "https://example.com/ktp.jpg",
		SelfiePhoto: "https://example.com/selfie.jpg",
	}, {
		ID:          uuid.New().String(),
		NIK:         random.String(16, random.Numeric),
		FullName:    "Annisa",
		LegalName:   "Annisa",
		BirthPlace:  "Bandung",
		BirthDate:   time.Now(),
		Salary:      9_000_000,
		KTPPhoto:    "https://example.com/ktp.jpg",
		SelfiePhoto: "https://example.com/selfie.jpg",
	}, {
		ID:          uuid.New().String(),
		NIK:         random.String(16, random.Numeric),
		FullName:    "Bagas",
		LegalName:   "Bagas",
		BirthPlace:  "Surabaya",
		BirthDate:   time.Now(),
		Salary:      10_500_000,
		KTPPhoto:    "https://example.com/ktp.jpg",
		SelfiePhoto: "https://example.com/selfie.jpg",
	},
}

var creditLimits = []entity.CreditLimit{
	{
		CustomerID:     customers[0].ID,
		TenorInMonth:   1,
		LimitAmount:    100_000,
		RemainingLimit: 100_000,
	}, {
		CustomerID:     customers[0].ID,
		TenorInMonth:   2,
		LimitAmount:    200_000,
		RemainingLimit: 200_000,
	}, {
		CustomerID:     customers[0].ID,
		TenorInMonth:   3,
		LimitAmount:    500_000,
		RemainingLimit: 500_000,
	}, {
		CustomerID:     customers[0].ID,
		TenorInMonth:   6,
		LimitAmount:    700_000,
		RemainingLimit: 700_000,
	}, {
		CustomerID:     customers[1].ID,
		TenorInMonth:   1,
		LimitAmount:    1_000_000,
		RemainingLimit: 1_000_000,
	}, {
		CustomerID:     customers[1].ID,
		TenorInMonth:   2,
		LimitAmount:    1_200_000,
		RemainingLimit: 1_200_000,
	}, {
		CustomerID:     customers[1].ID,
		TenorInMonth:   3,
		LimitAmount:    1_500_000,
		RemainingLimit: 1_500_000,
	}, {
		CustomerID:     customers[1].ID,
		TenorInMonth:   6,
		LimitAmount:    2_000_000,
		RemainingLimit: 2_000_000,
	}, {
		CustomerID:     customers[2].ID,
		TenorInMonth:   1,
		LimitAmount:    1_100_000,
		RemainingLimit: 1_100_000,
	}, {
		CustomerID:     customers[2].ID,
		TenorInMonth:   2,
		LimitAmount:    1_400_000,
		RemainingLimit: 1_400_000,
	}, {
		CustomerID:     customers[2].ID,
		TenorInMonth:   3,
		LimitAmount:    1_800_000,
		RemainingLimit: 1_800_000,
	}, {
		CustomerID:     customers[2].ID,
		TenorInMonth:   6,
		LimitAmount:    3_500_000,
		RemainingLimit: 3_500_000,
	},
}
