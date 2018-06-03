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
	id := c.Param("paymentTagId")
	var paymentTag models.PaymentTag
	db.Find(&paymentTag, id)

	if paymentTag.Id != 0 {
		return c.JSON(http.StatusOK, paymentTag)
	}
	err := fmt.Errorf("paymentTagId=%s is not found", id)
	return echo.NewHTTPError(http.StatusNotFound, err.Error())
}

func PostPaymentTag(c echo.Context) error {
	db := models.InitPaymentTagTable()
	defer db.Close()
	var paymentTag models.PaymentTag
	c.Bind(&paymentTag)
	fmt.Println(paymentTag)

	if paymentTag.PaymentId > 0 && paymentTag.TagId > 0 {
		db.Create(&paymentTag)
		return c.JSON(http.StatusCreated, paymentTag)
	}
	str := fmt.Sprintf("Values must be int %d: %d", paymentTag.PaymentId, paymentTag.TagId)
	err := errors.New(str)
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

func UpdatePaymentTag(c echo.Context) error {
	db := models.InitPaymentTagTable()
	defer db.Close()
	id := c.Param("paymentTagId")
	var paymentTag models.PaymentTag
	db.First(&paymentTag, id)

	if paymentTag.Id > 0 {
		var newPaymentTag models.PaymentTag
		c.Bind(&newPaymentTag)
		if newPaymentTag.PaymentId > 0 && newPaymentTag.TagId > 0 {
			result := models.PaymentTag{
				Id:        paymentTag.Id,
				PaymentId: newPaymentTag.PaymentId,
				TagId:     newPaymentTag.TagId,
			}
			db.Save(&result)
			return c.JSON(http.StatusOK, result)
		}
		str := fmt.Sprintf("Values must be int %d %d", newPaymentTag.PaymentId, newPaymentTag.TagId)
		err := errors.New(str)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := fmt.Errorf("paymentTagId=%s is not found", id)
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

func DeletePaymentTag(c echo.Context) error {
	db := models.InitPaymentTagTable()
	defer db.Close()
	id := c.Param("paymentTagId")
	var paymentTag models.PaymentTag
	db.First(&paymentTag, id)

	if paymentTag.Id > 0 {
		db.Delete(&paymentTag)
		return c.NoContent(http.StatusNoContent)
	}
	err := fmt.Errorf("paymentTagId=%s is not found", id)
	return echo.NewHTTPError(http.StatusNotFound, err.Error())
}
