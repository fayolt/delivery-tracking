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
	log.Println("processor.CreateProcessors - INFO - Closing jobs channel")
	close(jobChannel)
}

// startProcessor
func startProcessor(jobChannel chan Job, wg *sync.WaitGroup, pf processorFunc) {
	for job := range jobChannel {
		_, err := processJob(job, pf)
		if err != nil {
			log.Printf("processor.CreateProcessors - ERROR - %v", err)
			job.fails++
			go func() {
				// Check reenqueue condition
				log.Printf("processor.startProcessors - INFO - Pushing back %v into the jobs queue", job.data)
				time.Sleep(time.Duration(job.fails) * time.Second)
				jobChannel <- job
			}()
		}
	}
	wg.Done()
}

func processJob(job Job, pf processorFunc) (interface{}, error) {
	return pf(job.data)
}
