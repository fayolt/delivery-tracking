package processor

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProcessJob(t *testing.T) {
	idFunc := func(value interface{}) (interface{}, error) {
		return value, nil
	}
	job := NewJob(5)
	got, _ := processJob(*job, idFunc)
	assert.Equal(t, 5, got)
	assert.IsType(t, 5, got)
}

func TestReEnqueueJob(t *testing.T) {
	jobs := make(chan Job, 1)
	job := NewJob("test")
	reEnqueueJob(jobs, *job)
	assert.Equal(t, 1, len(jobs))
	got := <-jobs
	assert.Equal(t, "test", got.data)
	close(jobs)
}

func TestCreateProcessorsPool(t *testing.T) {
	idFunc := func(value interface{}) (interface{}, error) {
		return value, nil
	}
	jobs := make(chan Job)
	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- *NewJob(fmt.Sprintf("test-%d", i))
		}
		close(jobs)
	}()
	go CreateProcessorsPool(jobs, 2, idFunc)
	time.Sleep(1 * time.Second)

	value, ok := <-jobs

	assert.Equal(t, false, ok)
	assert.Equal(t, Job{data: interface{}(nil), fails: 0}, value)

}

func TestStartProcessorWithReEnqueue(t *testing.T) {
	errFunc := func(value interface{}) (interface{}, error) {
		return nil, errors.New("test")
	}
	jobs := make(chan Job, 1)
	jobs <- *NewJob("test")

	wg := &sync.WaitGroup{}

	go startProcessor(jobs, wg, errFunc)

	time.Sleep(300 * time.Millisecond)

	value := <-jobs

	assert.Equal(t, "test", value.data)
	assert.Less(t, 0, value.fails)
	jobs = nil

}
