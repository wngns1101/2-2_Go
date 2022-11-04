package magazine

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
	// HomeAddress Address
}

type Employee struct {
	Name   string
	Salary float64
	Address
	// HomeAddress Address
}

type Address struct {
	String     string
	City       string
	State      string
	PostalCode string
}
