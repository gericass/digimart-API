package scraper

import "time"

type Instrument struct {
	Category     int
	Name         string
	Price        int
	Condition    string
	Status       string
	URL          string
	RegisterDate time.Time
}
