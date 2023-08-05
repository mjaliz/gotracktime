package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mjaliz/gotracktime/internal/handlers"
)

func routes() *echo.Echo {
	e := echo.New()
	e.GET("/", handlers.Repo.Home)
	return e
}
