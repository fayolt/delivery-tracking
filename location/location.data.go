package location

import (
	"context"
	"log"
	"time"

	"github.com/fayolt/delivery-tracking/database"
)

func getLocations() ([]Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	results, err := database.DbConn.QueryContext(ctx, `SELECT id,
	longitude,
	latitude,
	driver_id
	FROM locations`)

	if err != nil {
		log.Printf("location.data.getLocations - ERROR - %v", err)
		return nil, err
	}

	defer results.Close()

	locations := make([]Location, 0)
	for results.Next() {
		var location Location
		results.Scan(
			&location.ID,
			&location.Longitude,
			&location.Latitude,
			&location.DriverID,
		)
		locations = append(locations, location)
	}
	return locations, nil
}

// InsertLocation writes a location record to the database
func InsertLocation(location Location) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	lastInsertID := 0
	// Using QueryRowContext coupled with Scan instead of ExecContext as a workaround to get LastInsertID
	// with postgres database driver lib/pq
	// Ref: https://stackoverflow.com/questions/33382981/go-how-to-get-last-insert-id-on-postgresql-with-namedexec
	err := database.DbConn.QueryRowContext(ctx, `INSERT INTO locations
	(longitude,
		latitude,
		driver_id) VALUES ($1, $2, $3) RETURNING id`,
		location.Longitude,
		location.Latitude,
		location.DriverID,
	).Scan(&lastInsertID)
	if err != nil {
		log.Printf("location.data.InsertLocation - ERROR - %v", err)
		return 0, err
	}
	log.Printf("location.data.InsertLocation - INFO - %+v written to database", location)
	return lastInsertID, nil
}
