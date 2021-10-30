package customer

import (
	"agent-allocation/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type api struct {
	service ICustomerService
}

func NewAPI(s ICustomerService) *api {
	return &api{s}
}

func (a *api) CreateQueue(c echo.Context) error {
	var customer domain.Customer
	err := c.Bind(&customer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message: ": err.Error()})
	}
	err = a.service.CreateQueue(c.Request().Context(), customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message: ": err.Error()})
	}
	return c.JSON(http.StatusOK, nil)
}
