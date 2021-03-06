package router

import (
	"github.com/labstack/echo"
	"github.com/lapis-zero09/tada-server/api/v1"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e1 := e.Group("/api/v1")
	e1.GET("/users", v1.UserIndex)
	e1.GET("/users/:userId", v1.UserShow)

	// payments
	e1.GET("/payments", v1.GetPayments)
	e1.GET("/payments/:paymentId", v1.GetPayment)
	e1.POST("/payments", v1.PostPayment)
	e1.PUT("/payments/:paymentId", v1.UpdatePayment)
	e1.DELETE("/payments/:paymentId", v1.DeletePayment)

	// tags
	e1.GET("/tags", v1.GetTags)
	e1.GET("/tags/:tagId", v1.GetTag)
	e1.POST("/tags", v1.PostTag)
	e1.PUT("/tags/:tagId", v1.UpdateTag)
	e1.DELETE("/tags/:tagId", v1.DeleteTag)

	// paymentTag
	e1.GET("/payment_tags", v1.GetPaymentTags)
	e1.GET("/payment_tags/:paymentTagId", v1.GetPaymentTag)
	e1.POST("/payment_tags", v1.PostPaymentTag)
	e1.PUT("/payment_tags/:paymentTagId", v1.UpdatePaymentTag)
	e1.DELETE("/payment_tags/:paymentTagId", v1.DeletePaymentTag)

	return e
}
