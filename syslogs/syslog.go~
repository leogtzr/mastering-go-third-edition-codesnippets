package main

import (
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_INFO, "systemLog.go")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
		log.Print("Everything is fine")
	}
}
