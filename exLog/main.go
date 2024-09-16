package main

import (
	"exLog/db"
	"exLog/manager/handlers"
	"log"
	"os"
)

func main() {
	logFile := logFile()
	defer logFile.Close()

	log.SetOutput(logFile)

	log.Println("<START>")
	log.Println("This is Logging!")

	db.Load()
	handlers.Handler()
	db.Load()
	log.Println("<END>")
}

func logFile() *os.File {
	f, err := os.OpenFile("LogFile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
