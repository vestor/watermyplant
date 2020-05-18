package log

import (
	"log"
	"os"
)

type logger struct {
	*log.Logger
}

var std = log.New(os.Stderr, "", log.LstdFlags | log.Lshortfile)

func Get() *logger {
	return &logger{
		Logger: std,
	}
}