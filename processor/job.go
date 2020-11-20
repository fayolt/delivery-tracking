package processor

// Job ...
type Job struct {
	data  interface{}
	fails int
}

// NewJob ...
func NewJob(data interface{}) *Job {
	return &Job{
		data: data,
	}
}
