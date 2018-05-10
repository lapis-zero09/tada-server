package router

import (
	"github.com/labstack/echo"
	"github.com/lapis-zero09/tada-server/api/v1"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e1 := e.Group("/api/v1")
	e1.GET("/users", v1.UserIndex)
	e1.GET("/users/:user_id", v1.UserShow)

	e1.GET("/payments", v1.GetPayments)
	e1.GET("/payments/:payment_id", v1.GetPayment)
	// e1.POST("/payments", v1.PostPayment)
	// e1.PUT("/payments/:payment_id", v1.UpdatePayment)
	// e1.DELETE("/payments/:payment_id", v1.DeletePayment)
	return e
}
