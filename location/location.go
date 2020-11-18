package location

// Location is a representation of the current location of a driver
//payload expcted by the http endpoint
type Location struct {
	DriverID  uint `json:"driver_id"`
	Latitude  int  `json:"latitude"`
	Longitude int  `json:"longitude"`
}

// { “latitude” : 123, “longitude” : 123, “driver_id”: 1 }
