package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/fayolt/delivery-tracking/database"
	"github.com/fayolt/delivery-tracking/location"
	"github.com/fayolt/delivery-tracking/processor"
	_ "github.com/lib/pq"
)

const apiBasePath = "/api/v1"

func main() {
	var (
		// maxQueueSize = flag.Int("max_queue_size", 100, "The size of job queue")
		maxProcessors = flag.Int("max_processors", 5, "The number of processors to start")
		port          = flag.String("port", "5000", "The server port")
	)

	flag.Parse()

	database.SetupDatabase()

	// Create jobs queue
	jobsQueue := make(chan processor.Job)

	registerControllers(apiBasePath, jobsQueue)

	// Create processor pool
	go processor.CreateProcessorsPool(jobsQueue, *maxProcessors, func(data interface{}) (interface{}, error) {
		return location.InsertLocation(data.(location.Location))
	})
	log.Printf("main - INFO - Starting server on port %s", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
