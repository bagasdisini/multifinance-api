package model

import "github.com/bagasdisini/multifinance-api/internal/entity"

type CustomerResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCustomerResponse(u *entity.Customer) *CustomerResponse {
	return &CustomerResponse{
		ID:   u.ID,
		Name: u.FullName,
	}
}

func NewCustomersResponse(u []entity.Customer) []CustomerResponse {
	cr := make([]CustomerResponse, len(u))
	for i, c := range u {
		cr[i] = CustomerResponse{
			ID:   c.ID,
			Name: c.FullName,
		}
	}
	return cr
}

type CustomerCreditLimitResponse struct {
	TenorInMonths  uint8   `json:"tenor_in_months"`
	LimitAmount    float64 `json:"limit_amount"`
	RemainingLimit float64 `json:"remaining_limit"`
}

func NewCustomerCreditLimitResponse(cr []entity.CreditLimit) *[]CustomerCreditLimitResponse {
	cl := make([]CustomerCreditLimitResponse, len(cr))
	for i, c := range cr {
		cl[i] = CustomerCreditLimitResponse{
			TenorInMonths:  c.TenorInMonth,
			LimitAmount:    c.LimitAmount,
			RemainingLimit: c.RemainingLimit,
		}
	}
	return &cl
}
