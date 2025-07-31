package handler

import (
	"context"
	"errors"
	"github.com/bagasdisini/multifinance-api/internal/entity"
	"github.com/bagasdisini/multifinance-api/internal/model"
	"github.com/bagasdisini/multifinance-api/internal/pkg/utils"
	"github.com/bagasdisini/multifinance-api/internal/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type TransactionHandler struct {
	db              *gorm.DB
	customerRepo    *repository.CustomerRepository
	creditLimitRepo *repository.CreditLimitRepository
	transactionRepo *repository.TransactionRepository
}

func NewTransactionHandler(e *echo.Echo, db *gorm.DB) *TransactionHandler {
	h := &TransactionHandler{
		db:              db,
		customerRepo:    repository.NewCustomerRepository(db),
		creditLimitRepo: repository.NewCreditLimitRepository(db),
		transactionRepo: repository.NewTransactionRepository(db),
	}

	e.GET("/api/transactions/:id", h.getTransactionByID)
	e.GET("/api/transactions/:customer_id", h.getTransactionsByCustomerID)

	e.POST("/api/transactions", h.createTransaction)

	return h
}

// Transaction
// @Tags Transactions
// @Summary Get Transaction
// @ID get-transaction
// @Param id path string true "Transaction ID"
// @Router /api/transactions/{id} [get]
// @Accept json
// @Produce json
// @Success 200
func (h *TransactionHandler) getTransactionByID(c echo.Context) error {
	id := c.Param("id")

	docs, err := h.transactionRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusBadRequest, model.ErrorResponse{
				Message: "Invalid Transaction ID",
			})
		}
		return model.NewInternalServerError()
	}

	return c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "Success",
		Data:    model.NewTransactionResponse(docs),
	})
}

// Transactions
// @Tags Transactions
// @Summary Get Transactions
// @ID get-transactions
// @Param customer_id path string true "Customer ID"
// @Router /api/transactions/{customer_id} [get]
// @Accept json
// @Produce json
// @Success 200
func (h *TransactionHandler) getTransactionsByCustomerID(c echo.Context) error {
	id := c.Param("customer_id")

	_, err := h.customerRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusBadRequest, model.ErrorResponse{
				Message: "Invalid Customer ID",
			})
		}
		return model.NewInternalServerError()
	}

	docs, err := h.transactionRepo.FindAllByCustomerID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusBadRequest, model.ErrorResponse{
				Message: "Invalid Customer ID",
			})
		}
		return model.NewInternalServerError()
	}

	return c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "Success",
		Data:    model.NewTransactionsResponse(docs),
	})
}

// Transaction
// @Tags Transactions
// @Summary Create Transaction
// @ID create-transaction
// @Param transaction body model.TransactionRequest true "Transaction Body"
// @Router /api/transactions [post]
// @Accept json
// @Produce json
// @Success 200
func (h *TransactionHandler) createTransaction(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 3*time.Second)
	defer cancel()
	c.SetRequest(c.Request().WithContext(ctx))

	f, err := model.NewTransactionRequest(c)
	if err != nil {
		return err
	}

	var response model.SuccessResponse
	err = h.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		customerRepo := repository.NewCustomerRepository(tx)
		creditRepo := repository.NewCreditLimitRepository(tx)
		transactionRepo := repository.NewTransactionRepository(tx)

		_, err := customerRepo.FindByID(f.CustomerID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return echo.NewHTTPError(http.StatusBadRequest, model.ErrorResponse{Message: "Invalid Customer ID"})
			}
			return model.NewInternalServerError()
		}

		cl, err := creditRepo.FindByCustomerIDAndTenorInMonthWithLock(f.CustomerID, f.TenorInMonth)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return echo.NewHTTPError(http.StatusBadRequest, model.ErrorResponse{Message: "Invalid Credit Limit"})
			}
			return model.NewInternalServerError()
		}

		totalValue := f.OTR + f.AdminFee + f.Interest
		transaction := entity.Transaction{
			CustomerID:   f.CustomerID,
			OTR:          f.OTR,
			AdminFee:     f.AdminFee,
			Interest:     f.Interest,
			TenorInMonth: f.TenorInMonth,
			AssetName:    f.AssetName,
			Status:       utils.TransactionStatusSuccess,
		}

		if totalValue > cl.RemainingLimit {
			transaction.Status = utils.TransactionStatusDeclined
			_ = transactionRepo.Create(transaction)
			return echo.NewHTTPError(http.StatusBadRequest, model.ErrorResponse{Message: "Insufficient credit limit"})
		}

		percentage := totalValue / cl.RemainingLimit

		if err := transactionRepo.Create(transaction); err != nil {
			return model.NewInternalServerError()
		}

		if err := creditRepo.UpdateCreditLimit(f.CustomerID, percentage); err != nil {
			return model.NewInternalServerError()
		}

		response = model.SuccessResponse{
			Message: "Success",
			Data:    model.NewTransactionResponse(transaction),
		}
		return nil
	})

	if err != nil {
		var httpErr *echo.HTTPError
		if errors.As(err, &httpErr) {
			return httpErr
		}
		return model.NewInternalServerError()
	}

	return c.JSON(http.StatusOK, response)
}
