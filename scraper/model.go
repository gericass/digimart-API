package scraper

import "time"

type Instrument struct {
	Category     string    `json:"category"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Price        int       `json:"price"`
	Condition    string    `json:"condition"`
	Status       string    `json:"status"`
	URL          string    `json:"url"`
	Image        string    `json:"image"`
	RegisterDate time.Time `json:"registerDate"`
}

type NewArrivalInstrument struct {
	Instrument
}