package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
		httpServer    *http.Server
	)

	flag.Parse()

	ctx := context.Background()

	database.InitDB(*dbHost, *dUser, *dbName, *dbPort)

	// Create jobs queue
	jobsQueue := make(chan processor.Job)

	router := registerControllers(apiBasePath, jobsQueue)

	// Create processors pool
	go processor.CreateProcessorsPool(jobsQueue, *maxProcessors, func(data interface{}) (interface{}, error) {
		return location.InsertLocation(data.(location.Location))
	})

	httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("main - INFO - Server started")

	<-done
	log.Println("main - INFO - Server shutting down")
	gracefullCtx, cancelShutdown := context.WithTimeout(ctx, 5*time.Second)
	defer func() {
		// Cleaning up
		close(jobsQueue)
		database.DbConn.Close()
		cancelShutdown()
	}()

	if err := httpServer.Shutdown(gracefullCtx); err != nil {
		log.Fatalf("main - ERROR - Server Shutdown Failed:%+v", err)
	}
	log.Println("main - INFO - Server Exited Properly")
}
