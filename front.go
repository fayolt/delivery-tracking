package main

import (
	"fmt"
	"log"

	"github.com/fayolt/delivery-tracking/location"
	"github.com/fayolt/delivery-tracking/processor"
	"github.com/gorilla/mux"
)

const locationsBasePath = "locations"

// registerControllers ...
func registerControllers(apiBasePath string, jobs chan processor.Job) *mux.Router {
	router := mux.NewRouter()
	lc := location.NewController(jobs)
	router.Handle(fmt.Sprintf("%s/%s", apiBasePath, locationsBasePath), *lc)
	log.Println("main.registerControllers - INFO - location controller successfully registred")
	return router
}
