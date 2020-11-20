package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fayolt/delivery-tracking/location"
	"github.com/fayolt/delivery-tracking/processor"
)

const locationsBasePath = "locations"

// registerControllers ...
func registerControllers(apiBasePath string, jobs chan processor.Job) {
	lc := location.NewController(jobs)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, locationsBasePath), *lc)
	log.Println("main.registerControllers - INFO - location controller successfully registred")
}
