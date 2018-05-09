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
	return e
}
