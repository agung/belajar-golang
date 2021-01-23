package customer

// Customer is model data
type Customer struct {
	name, address, position string
	age                     int
	salary                  float64
}

var customer Customer

// SetName is pass value to struct
func SetName(name string) {
	customer.name = name
}

// GetName is getting value from struct
func GetName() string {
	return customer.name
}
