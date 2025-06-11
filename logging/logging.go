package logging

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)

func LogError(err error) {
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}
}
