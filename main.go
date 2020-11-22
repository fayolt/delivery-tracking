package main

import (
	"flag"
	"fmt"
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
		dbHost        = flag.String("db_host", "localhost", "The database hostname")
		dUser         = flag.String("db_user", "postgres", "The database user")
		dbName        = flag.String("db_name", "delivery-tracking", "The name of the database")
		dbPort        = flag.Int("db_port", 5432, "The database port")
		maxProcessors = flag.Int("max_processors", 5, "The number of processors to start")
		port          = flag.Int("port", 5000, "The server port")
	)

	flag.Parse()

	database.InitDB(*dbHost, *dUser, *dbName, *dbPort)

	// Create jobs queue
	jobsQueue := make(chan processor.Job)

	registerControllers(apiBasePath, jobsQueue)

	// Create processor pool
	go processor.CreateProcessorsPool(jobsQueue, *maxProcessors, func(data interface{}) (interface{}, error) {
		return location.InsertLocation(data.(location.Location))
	})
	log.Printf("main - INFO - Starting server on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
