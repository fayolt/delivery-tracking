package location

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fayolt/delivery-tracking/processor"
)

// Controller ...
type Controller struct {
	jobs chan processor.Job
}

// NewController creates an instance of locationController
func NewController(jobChannel chan processor.Job) *Controller {
	fmt.Println("controller created successfull")
	return &Controller{
		jobs: jobChannel,
	}
}

func (controller Controller) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	locationsHandler(controller.jobs, rw, req)

}

func locationsHandler(jobs chan processor.Job, rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		// Add a new location
		var newLocation Location
		err := json.NewDecoder(req.Body).Decode(&newLocation)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		// Create Job and push the work onto the job channel.
		job := processor.NewJob(newLocation)
		jobs <- *job
		// go func() {

		// }()

		// _, err = insertLocation(newLocation)
		// if err != nil {
		// 	rw.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }
		rw.WriteHeader(http.StatusAccepted)
		return

	case http.MethodGet:
		fmt.Println("GET request received")
		locationList, err := getLocations()
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		locationsJSON, err := json.Marshal(locationList)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(locationsJSON)
		return
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
