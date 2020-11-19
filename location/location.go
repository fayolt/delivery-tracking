package location

// Location is a representation of the current location of a driver
//payload expcted by the http endpoint
type Location struct {
	ID        int `json:"id"`
	DriverID  int `json:"driver_id"`
	Latitude  int `json:"latitude"`
	Longitude int `json:"longitude"`
}

// { “latitude” : 123, “longitude” : 123, “driver_id”: 1 }

// Driver is representation of a driver's details
type Driver struct {
	ID        int
	firstName string
	lastName  string
}
