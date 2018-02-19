package Infrastructure

import "time"

// Instrument : Entity
type Instrument struct {
	ID int64
	Category int
	Name string
	Price int
	Condition string
	Status string
	URL string
	registerDate time.Time
}