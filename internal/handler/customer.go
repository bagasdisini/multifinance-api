package handler

import (
	"errors"
	"github.com/bagasdisini/multifinance-api/internal/model"
	"github.com/bagasdisini/multifinance-api/internal/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type CustomerHandler struct {
	customerRepo    *repository.CustomerRepository
	creditLimitRepo *repository.CreditLimitRepository
}

func NewCustomerHandler(e *echo.Echo, db *gorm.DB) *CustomerHandler {
	h := &CustomerHandler{
		customerRepo:    repository.NewCustomerRepository(db),
		creditLimitRepo: repository.NewCreditLimitRepository(db),
	}

	e.GET("/api/customers", h.getCustomers)
	e.GET("/api/customers/:id", h.getCustomerByID)
	e.GET("/api/customers/:id/credit-limit", h.getCreditLimitByCustomerID)

	return h
}

// Customers
// @Tags Customers
// @Summary Get Customers
// @ID get-customers
// @Router /api/customers [get]
// @Accept json
// @Produce json
// @Success 200
func (h *CustomerHandler) getCustomers(c echo.Context) error {
	docs, err := h.customerRepo.FindAll()
	if err != nil {
		return model.NewInternalServerError()
	}
	return c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "Success",
		Data:    model.NewCustomersResponse(docs),
	})
}

// Customer
// @Tags Customers
// @Summary Get Customer
// @ID get-customer
// @Param id path string true "Customer ID"
// @Router /api/customers/{id} [get]
// @Accept json
// @Produce json
// @Success 200
func (h *CustomerHandler) getCustomerByID(c echo.Context) error {
	id := c.Param("id")

	docs, err := h.customerRepo.FindByID(id)
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
		Data:    model.NewCustomerResponse(docs),
	})
}

// Credit Limit
// @Tags Customers
// @Summary Get Credit Limit
// @ID get-credit-limit
// @Param id path string true "Customer ID"
// @Router /api/customers/{id}/credit-limit [get]
// @Accept json
// @Produce json
// @Success 200
func (h *CustomerHandler) getCreditLimitByCustomerID(c echo.Context) error {
	id := c.Param("id")

	_, err := h.customerRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusBadRequest, model.ErrorResponse{
				Message: "Invalid Customer ID",
			})
		}
		return model.NewInternalServerError()
	}

	docs, err := h.creditLimitRepo.FindAllByCustomerID(id)
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
		Data:    model.NewCustomerCreditLimitResponse(docs),
	})
}
