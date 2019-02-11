package models

import (
	"database/sql"
	"os"
	"util"
)

var db *sql.DB

// Init is to initialize connection object.
func Init() {
	var err error

	db, err = sql.Open("postgres", os.Getenv("CONNECTION_STRING"))
	util.CheckError(err)
}
