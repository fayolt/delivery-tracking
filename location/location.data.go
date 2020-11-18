package location

import (
	"context"
	"time"

	"github.com/fayolt/delivery-tracking/database"
)

func getLocations() ([]Location, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := "SELECT FROM locations"
	results, err := database.DbConn.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer results.Close()

	locations := make([]Location, 0)
	for results.Next() {
		var location Location
		results.Scan()
		locations = append(locations, location)
	}
	return locations, nil
}

func insertLocation(location Location) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := "INSERT INTO locations () VALUES "
	result, err := database.DbConn.ExecContext(ctx, query)
	if err != nil {
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(insertID), nil
}
