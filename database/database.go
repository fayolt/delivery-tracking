package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "delivery-tracking"
)

// DbConn a database connection
var DbConn *sql.DB

// SetupDatabase ...
func SetupDatabase() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	DbConn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetMaxOpenConns(5)
	DbConn.SetMaxIdleConns(5)
	DbConn.SetConnMaxLifetime(60 * time.Second)
	// defer DbConn.Close()
}
