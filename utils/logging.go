package utils

import (
	"log"
	"os"
	"time"
)

func SetupLogging() {
	fileName := "storage/log-" + time.Now().Format("2006-01-02") + ".log"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	log.SetOutput(file)
}
