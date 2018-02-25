package main

import (
	"github.com/labstack/echo"
	"github.com/gericass/digimart-API/handler"
)

func main() {
	e := echo.New()
	e.GET("/search", handler.SearchInstrumentsHandler)
}
