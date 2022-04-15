package customer

// Customer struct
type Customer struct {
	Name     string
	Email    string
	Password string `json:"-"`
}
