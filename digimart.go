package main

import (
	"github.com/labstack/echo"
	"github.com/gericass/digimart-API/handler"
	"github.com/gericass/digimart-API/Infrastructure"
)

func dbMiddleware(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := Infrastructure.ConnectDB()
		if err != nil {
			return err
		}
		cc := &handler.CustomContext{c, db}
		return h(cc)
	}
}

func main() {
	e := echo.New()
	e.Use(dbMiddleware)

	e.GET("/search", handler.SearchInstrumentsHandler)
}
