package Infrastructure

import "time"

// Instrument : Entity
type Instrument struct {
	ID           int64
	Category     int
	Name         string
	Price        int
	Condition    string
	Status       string
	URL          string
	RegisterDate time.Time
}

// User : Entity
type User struct {
	ID       int64
	Name     string
	Password string
	Salt     string
}
