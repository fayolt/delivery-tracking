package location

import (
	"context"
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

func insertLocation(location Location) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := database.DbConn.ExecContext(ctx, `INSERT INTO locations
	(longitude,
		latitude,
		driver_id) VALUES ($1, $2, $3)`,
		location.Longitude,
		location.Latitude,
		location.DriverID,
	)
	if err != nil {
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(insertID), nil
}
