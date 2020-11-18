package location

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const locationsBasePath = "locations"

// SetupRoutes ...
func SetupRoutes(apiBasePath string) {
	http.HandleFunc(fmt.Sprintf("%s/%s", apiBasePath, locationsBasePath), locationHandler)
}

func locationHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		// Add a new location
		var newLocation Location
		err := json.NewDecoder(req.Body).Decode(&newLocation)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// Check if driver_id exists
		_, err = addLocation(newLocation)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(http.StatusCreated)
		return

	case http.MethodGet:
		locationList := getLocations()
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
