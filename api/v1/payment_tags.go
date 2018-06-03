package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lapis-zero09/tada-server/models"
)

func intInSlice(num int, list []int) bool {
	for _, numInList := range list {
		if numInList == num {
			return true
		}
	}
	return false
}

func GetPaymentTags(c echo.Context) error {
	db := models.InitPaymentTable()
	db = models.InitTagTable()
	db = models.InitPaymentTagTable()
	defer db.Close()

	var paymentTags []models.PaymentTag
	err := db.Joins("JOIN payments ON payment_tags.payment_id = payments.id").
		Joins("JOIN tags ON payment_tags.tag_id = tags.id").
		Preload("Payment").
		Preload("Tag").
		Find(&paymentTags).
		Error
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, paymentTags)
}

func GetPaymentTag(c echo.Context) error {
	db := models.InitPaymentTable()
	db = models.InitTagTable()
	db = models.InitPaymentTagTable()
	defer db.Close()

	id := c.Param("paymentTagId")
	var paymentTag models.PaymentTag
	db.Joins("JOIN payments ON payment_tags.payment_id = payments.id").
		Joins("JOIN tags ON payment_tags.tag_id = tags.id").
		Preload("Payment").
		Preload("Tag").
		First(&paymentTag, id)

	if paymentTag.ID > 0 {
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

	var paymentIds []int
	db.Model(&models.Payment{}).Pluck("id", &paymentIds)
	var tagIds []int
	db.Model(&models.Tag{}).Pluck("id", &tagIds)

	if intInSlice(paymentTag.PaymentID, paymentIds) && intInSlice(paymentTag.TagID, tagIds) {
		db.Create(&paymentTag)
		return c.JSON(http.StatusCreated, paymentTag)
	}
	str := fmt.Sprintf("Values must be exsists %d: %d", paymentTag.PaymentID, paymentTag.TagID)
	err := errors.New(str)
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

func UpdatePaymentTag(c echo.Context) error {
	db := models.InitPaymentTagTable()
	defer db.Close()

	id := c.Param("paymentTagId")
	var paymentTag models.PaymentTag
	db.First(&paymentTag, id)

	var paymentIds []int
	db.Model(&models.Payment{}).Pluck("id", &paymentIds)
	fmt.Println(paymentIds)
	var tagIds []int
	db.Model(&models.Tag{}).Pluck("id", &tagIds)
	fmt.Println(tagIds)

	if paymentTag.ID > 0 {
		var newPaymentTag models.PaymentTag
		c.Bind(&newPaymentTag)
		if intInSlice(newPaymentTag.PaymentID, paymentIds) && intInSlice(newPaymentTag.TagID, tagIds) {
			result := models.PaymentTag{
				ID:        paymentTag.ID,
				PaymentID: newPaymentTag.PaymentID,
				TagID:     newPaymentTag.TagID,
			}
			db.Save(&result)
			return c.JSON(http.StatusOK, result)
		}
		str := fmt.Sprintf("Values must be exists %d %d", newPaymentTag.PaymentID, newPaymentTag.TagID)
		err := errors.New(str)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := fmt.Errorf("paymentTagID=%s is not found", id)
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

func DeletePaymentTag(c echo.Context) error {
	db := models.InitPaymentTagTable()
	defer db.Close()

	id := c.Param("paymentTagId")
	var paymentTag models.PaymentTag
	db.First(&paymentTag, id)

	if paymentTag.ID > 0 {
		db.Delete(&paymentTag)
		return c.NoContent(http.StatusNoContent)
	}
	err := fmt.Errorf("paymentTagID=%s is not found", id)
	return echo.NewHTTPError(http.StatusNotFound, err.Error())
}
