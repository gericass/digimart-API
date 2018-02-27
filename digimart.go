package main

import (
	"github.com/labstack/echo"
	"github.com/gericass/digimart-API/handler"
	"github.com/gericass/digimart-API/infrastructure"
)

func dbMiddleware(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := infrastructure.ConnectDB()
		if err != nil {
			return err
		}
		defer db.Close()
		cc := &handler.CustomContext{c, db}
		return h(cc)
	}
}

func main() {
	e := echo.New()
	e.Use(dbMiddleware)

	e.GET("/", handler.NewArrivalHandler)
	e.GET("/search", handler.SearchInstrumentsHandler)
	e.POST("/user/register", handler.RegisterUserHandler)
	e.POST("/user/login", handler.LoginHandler)
	e.POST("/user/logout", handler.LogoutHandler)
	e.GET("/user/subscribe", handler.GetSubscribeInstrumentsHandler)
	e.POST("/user/subscribe", handler.SubscribeInstrumentHandler)
	e.DELETE("/user/unsubscribe", handler.UnSubscribeInstrumentHandler)

	e.Start(":8000")
}
