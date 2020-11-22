package processor

import (
	"log"
	"sync"
	"time"
)

type processorFunc func(interface{}) (interface{}, error)

// CreateProcessorsPool creates a pool of poolSize processors
func CreateProcessorsPool(jobChannel chan Job, poolSize int, pf processorFunc) {
	var wg sync.WaitGroup
	for i := 1; i <= poolSize; i++ {
		wg.Add(1)
		log.Printf("processor.CreateProcessors - INFO - Starting processor %d", i)

		go startProcessor(jobChannel, &wg, pf)

	}
	wg.Wait()
}

// startProcessor
func startProcessor(jobChannel chan Job, wg *sync.WaitGroup, pf processorFunc) {
	for job := range jobChannel {
		_, err := processJob(job, pf)
		if err != nil {
			log.Printf("processor.CreateProcessors - ERROR - %v", err)
			job.fails++
			go reEnqueueJob(jobChannel, job)
		}
	}
	log.Println("processor.startProcessor - INFO - processor done")
	wg.Done()
}

func processJob(job Job, pf processorFunc) (interface{}, error) {
	return pf(job.data)
}

func reEnqueueJob(jobs chan Job, job Job) {
	// Check reenqueue conditions???
	log.Printf("processor.reEnqueueJob - INFO - Pushing back %+v into the jobs queue", job.data)
	time.Sleep(time.Duration(job.fails) * time.Second)
	jobs <- job
}
