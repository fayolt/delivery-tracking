package processor

import (
	"fmt"
	"sync"
	"time"
)

type processorFunc = func(interface{}) (interface{}, error)

// CreateProcessorsPool ...
func CreateProcessorsPool(jobChannel chan Job, poolSize int, pf processorFunc) {
	var wg sync.WaitGroup
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go createProcessor(jobChannel, &wg, pf)

	}
	wg.Wait()
	close(jobChannel)
}

// createProcessor
func createProcessor(jobChannel chan Job, wg *sync.WaitGroup, pf processorFunc) {
	fmt.Println("inside worker")
	for job := range jobChannel {
		_, err := processJob(job, pf)
		if err != nil {
			job.fails++
			go func() {
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
