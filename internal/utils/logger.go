package utils

import (
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	Logger = log.New(os.Stdout, "BugFlow: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogInfo(message string) {
	Logger.Println("INFO:", message)
}

func LogError(message string, err error) {
	Logger.Printf("ERROR: %s - %v\n", message, err)
}

func LogFatal(message string, err error) {
	Logger.Fatalf("FATAL: %s - %v\n", message, err)
}
