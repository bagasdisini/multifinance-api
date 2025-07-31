package model

import (
	"github.com/bagasdisini/multifinance-api/internal/entity"
	"github.com/bagasdisini/multifinance-api/internal/pkg/validate"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TransactionRequest struct {
	CustomerID   string  `json:"customer_id" form:"customer_id"`
	OTR          float64 `json:"otr" form:"otr"`
	AdminFee     float64 `json:"admin_fee" form:"admin_fee"`
	Interest     float64 `json:"interest" form:"interest"`
	TenorInMonth uint8   `json:"tenor_in_month" form:"tenor_in_month"`
	AssetName    string  `json:"asset_name" form:"asset_name"`
}

func NewTransactionRequest(c echo.Context) (*TransactionRequest, error) {
	form := new(TransactionRequest)
	if err := c.Bind(&form); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{
			Message: "Invalid form",
		})
	}

	validator := validate.Validator{}

	validator.Check("customer_id", form.CustomerID, validate.EmptyStringRule)
	validator.Check("otr", form.OTR, validate.PositiveFloatRule)
	validator.Check("admin_fee", form.AdminFee, validate.PositiveFloatRule)
	validator.Check("interest", form.Interest, validate.PositiveFloatRule)
	validator.Check("tenor_in_month", form.TenorInMonth, validate.PositiveUint8Rule)
	validator.Check("asset_name", form.AssetName, validate.EmptyStringRule)

	if len(validator.Errors) > 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{
			Message: "Validation failed",
			Errors:  validator.Errors,
		})
	}

	return form, nil
}

type TransactionResponse struct {
	ID           string  `json:"id"`
	OTR          float64 `json:"otr"`
	AdminFee     float64 `json:"admin_fee"`
	Interest     float64 `json:"interest"`
	TenorInMonth uint8   `json:"tenor_in_month"`
	AssetName    string  `json:"asset_name"`
	Status       string  `json:"status"`
}

func NewTransactionResponse(u entity.Transaction) *TransactionResponse {
	return &TransactionResponse{
		ID:           u.ID,
		OTR:          u.OTR,
		AdminFee:     u.AdminFee,
		Interest:     u.Interest,
		TenorInMonth: u.TenorInMonth,
		AssetName:    u.AssetName,
		Status:       u.Status,
	}
}

func NewTransactionsResponse(u []entity.Transaction) []TransactionResponse {
	tr := make([]TransactionResponse, len(u))
	for i, t := range u {
		tr[i] = TransactionResponse{
			ID:           t.ID,
			OTR:          t.OTR,
			AdminFee:     t.AdminFee,
			Interest:     t.Interest,
			TenorInMonth: t.TenorInMonth,
			AssetName:    t.AssetName,
			Status:       t.Status,
		}
	}
	return tr
}
