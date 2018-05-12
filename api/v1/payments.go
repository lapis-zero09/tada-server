package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lapis-zero09/tada-server/models"
)

func GetPayments(c echo.Context) error {
	db := models.InitDb()
	defer db.Close()

	var payments []models.Payment
	// SELECT * FROM users
	db.Find(&payments)

	return c.JSON(http.StatusOK, payments)
}

func GetPayment(c echo.Context) error {
	db := models.InitDb()
	defer db.Close()
	id := c.Param("payment_id")
	var payment models.Payment
	db.First(&payment, id)

	if payment.Id != 0 {
		return c.JSON(http.StatusOK, payment)
	} else {
		err := fmt.Errorf("payment_id=%d is not found", id)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
}

func PostPayment(c echo.Context) error {
	db := models.InitDb()
	defer db.Close()

	var payment models.Payment
	c.Bind(&payment)

	db.Create(&payment)
	return c.JSON(http.StatusOK, payment)
	// err := errors.New("Fields are empty")
	// return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

func UpdatePayment(c echo.Context) error {
	db := models.InitDb()
	defer db.Close()
	id := c.Param("payment_id")
	var payment models.Payment
	db.First(&payment, id)

	if payment.Id != 0 {
		var newPayment models.Payment
		c.Bind(&newPayment)

		result := models.Payment{
			Id:      payment.Id,
			PlaceId: newPayment.PlaceId,
			Cost:    newPayment.Cost,
		}

		db.Save(&result)
		return c.JSON(http.StatusOK, newPayment)
	} else {
		err := errors.New("User not found")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// err := errors.New("Fields are empty")
	// return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

func DeletePayment(c echo.Context) error {
	db := models.InitDb()
	defer db.Close()
	id := c.Param("payment_id")
	var payment models.Payment
	db.First(&payment, id)

	if payment.Id != 0 {
		db.Delete(&payment)
		return c.NoContent(http.StatusNoContent)
	} else {
		err := fmt.Errorf("payment_id=%d is not found", id)
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
}
