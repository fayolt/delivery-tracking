package main

import (
	"net/http"

	"github.com/fayolt/delivery-tracking/database"
	"github.com/fayolt/delivery-tracking/location"
	_ "github.com/lib/pq"
)

const apiBasePath = "/api/v1"

func main() {
	database.SetupDatabase()
	location.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}
