package db

import (
	_ "github.com/joho/godotenv/autoload"
	scribble "github.com/nanobox-io/golang-scribble"
	logger "github.com/vestor/watermyplant/internal/log"
	"os"
)

var l = logger.Get()
var database = &scribble.Driver{}
var dbSetup = false

func SetupDB() {
	dir := os.Getenv("DB_DIR")
	l.Printf("Setting up DB in %v\n", dir)
	var db, dbErr = scribble.New(dir, nil)
	if dbErr != nil {
		l.Println("Error", dbErr)
	}
	dbSetup = true
	database = db
}

func GetDb() *scribble.Driver {
	if !dbSetup {
		SetupDB()
	}
	return database
}
