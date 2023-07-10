package logs

import (
	"log"
	"os"
)

const (
	Fatal = "FATAL"
	Info  = "INFO"
	Error = "ERROR"
)

var InfoLog *log.Logger
var ErrorLog *log.Logger
var FatalLog *log.Logger

func SetUpLog() *os.File {
	file, err := os.OpenFile("logs/logs.txt", os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalf("Failed to open log file %s", err.Error())
	}
	InfoLog = log.New(file, "INFO\t", log.Ldate|log.Ltime)
	ErrorLog = log.New(file, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	FatalLog = log.New(file, "FATAL\t", log.Ldate|log.Ltime|log.Lshortfile)
	return file
}

func Log(logType, message, method, urlPath string, userID int) {
	switch logType {

	case "ERROR":
		if userID == 0 {
			ErrorLog.Printf("%s %s %s\n", method, urlPath, message)
			return
		}
		ErrorLog.Printf("%d %s %s %s\n", userID, method, urlPath, message)
	case "FATAL":
		FatalLog.Fatalf("%d %s %s %s\n", userID, method, urlPath, message)
	case "INFO":
		InfoLog.Printf("%d %s %s\n", userID, method, urlPath)
	default:
		log.Printf("%d %s %s %s\n", userID, method, urlPath, message)
	}
}
