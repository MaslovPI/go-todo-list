package logging

import (
	"log"
	"os"
)

var (
	errLogger  = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
	infoLogger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
)

func LogError(err error) {
	if err != nil {
		errLogger.Println(err)
	}
}

func LogFatal(err error) {
	LogError(err)
	if err != nil {
		os.Exit(1)
	}
}

func LogInfo(message string) {
	infoLogger.Panicln(message)
}
