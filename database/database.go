package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// DbConn a database connection
var DbConn *sql.DB

// SetupDatabase creates a connection pool to the database
func SetupDatabase(host, dbUser, dbName string, port int) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, dbUser, dbName)
	DbConn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetMaxOpenConns(5)
	DbConn.SetMaxIdleConns(5)
	DbConn.SetConnMaxLifetime(60 * time.Second)
	log.Println("database.SetupDatabase - Info - Connection to database established")

	// defer DbConn.Close()
}
