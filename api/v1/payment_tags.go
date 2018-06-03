package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lapis-zero09/tada-server/models"
)

func GetPaymentTags(c echo.Context) error {
	db := models.InitPaymentTagTable()
	defer db.Close()
	var paymentTags []models.PaymentTag
	db.Find(&paymentTags)

	return c.JSON(http.StatusOK, paymentTags)
}

func GetPaymentTag(c echo.Context) error {
	db := models.InitPaymentTagTable()
	defer db.Close()
	id := c.Param("paymentTagID")
	var paymentTag models.PaymentTag
	db.Find(&paymentTag, id)

	if paymentTag.ID > 0 {
		return c.JSON(http.StatusOK, paymentTag)
	}
	err := fmt.Errorf("paymentTagID=%s is not found", id)
	return echo.NewHTTPError(http.StatusNotFound, err.Error())
}

func PostPaymentTag(c echo.Context) error {
	db := models.InitPaymentTagTable()
	defer db.Close()
	var paymentTag models.PaymentTag
	c.Bind(&paymentTag)
	fmt.Println(paymentTag)

	if paymentTag.PaymentID > 0 && paymentTag.TagID > 0 {
		db.Create(&paymentTag)
		return c.JSON(http.StatusCreated, paymentTag)
	}
	str := fmt.Sprintf("Values must be int %d: %d", paymentTag.PaymentID, paymentTag.TagID)
	err := errors.New(str)
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

func UpdatePaymentTag(c echo.Context) error {
	db := models.InitPaymentTagTable()
	defer db.Close()
	id := c.Param("paymentTagID")
	var paymentTag models.PaymentTag
	db.First(&paymentTag, id)

	if paymentTag.ID > 0 {
		var newPaymentTag models.PaymentTag
		c.Bind(&newPaymentTag)
		if newPaymentTag.PaymentID > 0 && newPaymentTag.TagID > 0 {
			result := models.PaymentTag{
				ID:        paymentTag.ID,
				PaymentID: newPaymentTag.PaymentID,
				TagID:     newPaymentTag.TagID,
			}
			db.Save(&result)
			return c.JSON(http.StatusOK, result)
		}
		str := fmt.Sprintf("Values must be int %d %d", newPaymentTag.PaymentID, newPaymentTag.TagID)
		err := errors.New(str)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := fmt.Errorf("paymentTagID=%s is not found", id)
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

func DeletePaymentTag(c echo.Context) error {
	db := models.InitPaymentTagTable()
	defer db.Close()
	id := c.Param("paymentTagID")
	var paymentTag models.PaymentTag
	db.First(&paymentTag, id)

	if paymentTag.ID > 0 {
		db.Delete(&paymentTag)
		return c.NoContent(http.StatusNoContent)
	}
	err := fmt.Errorf("paymentTagID=%s is not found", id)
	return echo.NewHTTPError(http.StatusNotFound, err.Error())
}
