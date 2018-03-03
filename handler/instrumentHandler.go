package handler

import (
	"github.com/labstack/echo"
	"github.com/gericass/digimart-API/scraper"
	"encoding/json"
	"strconv"
)

func SearchInstrumentsHandler(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return err
	}
	insts, err := scraper.Scrape(keyword, page)
	if err != nil {
		return err
	}
	j, _ := json.Marshal(insts)
	return c.String(200, string(j))
}

func NewArrivalHandler(c echo.Context) error {
	insts, err := scraper.NewArrival()
	if err != nil {
		return err
	}
	instsJson, err := json.Marshal(insts)
	if err != nil {
		return err
	}
	return c.String(200, string(instsJson))
}
