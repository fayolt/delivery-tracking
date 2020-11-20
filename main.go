package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fayolt/delivery-tracking/database"
	"github.com/fayolt/delivery-tracking/location"
	"github.com/fayolt/delivery-tracking/processor"
	_ "github.com/lib/pq"
)

const apiBasePath = "/api/v1"

const locationsBasePath = "locations"

// SetupRoutes ...
func SetupRoutes(apiBasePath string, jobs chan processor.Job) {
	lc := location.NewController(jobs)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, locationsBasePath), *lc)
	fmt.Println("route setup")
}

func main() {
	// var (
	// 	maxQueueSize = flag.Int("max_queue_size", 100, "The size of job queue")
	// 	maxWorkers   = flag.Int("max_workers", 5, "The number of processors to start")
	// 	port         = flag.String("port", "5000", "The server port")
	// )

	// maxQueueSize := 100
	maxWorkers := 5
	port := "5000"
	// flag.Parse()

	database.SetupDatabase()

	// Create jobs queue
	jobsQueue := make(chan processor.Job)

	SetupRoutes(apiBasePath, jobsQueue)

	// Create processor pool
	processor.CreateProcessorsPool(jobsQueue, maxWorkers, func(data interface{}) (interface{}, error) {
		return location.InsertLocation(data.(location.Location))
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
