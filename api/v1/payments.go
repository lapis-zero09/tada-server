package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lapis-zero09/tada-server/models"
)

func GetPayments(c echo.Context) error {
	db := models.InitPaymentTable()
	defer db.Close()
	var payments []models.Payment
	db.Find(&payments)

	return c.JSON(http.StatusOK, payments)
}

func GetPayment(c echo.Context) error {
	db := models.InitPaymentTable()
	defer db.Close()
	id := c.Param("paymentId")
	var payment models.Payment
	db.First(&payment, id)

	if payment.ID > 0 {
		return c.JSON(http.StatusOK, payment)
	}
	err := fmt.Errorf("paymentId=%s is not found", id)
	return echo.NewHTTPError(http.StatusNotFound, err.Error())
}

func PostPayment(c echo.Context) error {
	db := models.InitPaymentTable()
	defer db.Close()
	var payment models.Payment
	c.Bind(&payment)

	if payment.PlaceID > 0 && payment.Cost > 0 {
		db.Create(&payment)
		return c.JSON(http.StatusCreated, payment)
	}
	err := errors.New("Values must be int")
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

func UpdatePayment(c echo.Context) error {
	db := models.InitPaymentTable()
	defer db.Close()
	id := c.Param("paymentId")
	var payment models.Payment
	db.First(&payment, id)

	if payment.ID > 0 {
		var newPayment models.Payment
		c.Bind(&newPayment)
		if newPayment.PlaceID > 0 && newPayment.Cost > 0 {
			result := models.Payment{
				ID:      payment.ID,
				PlaceID: newPayment.PlaceID,
				Cost:    newPayment.Cost,
			}
			db.Save(&result)
			return c.JSON(http.StatusOK, result)
		}
		err := errors.New("Values must be int")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := fmt.Errorf("paymentId=%s is not found", id)
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

func DeletePayment(c echo.Context) error {
	db := models.InitPaymentTable()
	defer db.Close()
	id := c.Param("paymentId")
	var payment models.Payment
	db.First(&payment, id)

	if payment.ID > 0 {
		db.Delete(&payment)
		return c.NoContent(http.StatusNoContent)
	}
	err := fmt.Errorf("paymentId=%s is not found", id)
	return echo.NewHTTPError(http.StatusNotFound, err.Error())
}
