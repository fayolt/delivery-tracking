package processor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewJob tests the creation of a job instance
func TestNewJob(t *testing.T) {
	var data interface{}
	assert.IsType(t, NewJob(data), &Job{})
}
