package Infrastructure

// User : Entity
type User struct {
	ID       int64
	Name     string
	Password string
	Salt     string
}
