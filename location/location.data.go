package location

import "sync"

var locationMap = struct {
	sync.RWMutex
	m map[int]Location
}{m: make(map[int]Location)}

var index = 0

func getLocations() []Location {
	locationMap.RLock()
	locations := make([]Location, 0, len(locationMap.m))
	for _, value := range locationMap.m {
		locations = append(locations, value)
	}
	locationMap.RUnlock()
	return locations
}

func addLocation(location Location) (int, error) {
	index++
	locationMap.Lock()
	locationMap.m[index] = location
	locationMap.Unlock()

	return index, nil
}
