package sqlite

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Db is our database struct used for interacting with the database
var Db *sql.DB

// New makes a new database using the connection string and
// returns it, otherwise returns the error
func InitDB(dbPath string) {
	// open database
	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		log.Panic(err)
	}

	// Check that our connection is good
	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}
	Db = db
}

func CloseDB() error {
	return Db.Close()
}
