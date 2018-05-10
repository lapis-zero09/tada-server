package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/lapis-zero09/tada-server/models"
)

// import "github.com/labstack/echo"

func GetPayments(c echo.Context) error {
	return c.JSON(http.StatusOK, models.SamplePayments())
}

func GetPayment(c echo.Context) error {
	payments := models.SamplePayments()
	id, err := strconv.Atoi(c.Param("payment_id"))
	if err != nil {
		return err
	}
	if id > len(payments)-1 {
		err := fmt.Errorf("payment_id=%d is not found", id)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, payments[id])
}

// func PostPayment(c echo.Context) error {}

// func UpdataPayment(c echo.Context) error {}

// func DeletePayment(c echo.Context) error {}
