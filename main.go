package main

import (
	"net/http"

	"github.com/fayolt/delivery-tracking/location"
)

const apiBasePath = "/api/v1"

func main() {
	location.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}
