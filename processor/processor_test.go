package processor

import (
	"testing"

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
}
