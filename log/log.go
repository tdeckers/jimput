package log

import (
	"log"
	"os"
	"time"
)

var (
	tag = string
)

func init() {
	tag = os.Args[0]
}

func log(level, msg string) {
	var w *os.File
	timestamp := time.Now().Format(time.RFC3339)
	hostname, _ := os.Hostname()
	switch level {
	case "DEBUG", "INFO", "NOTICE", "WARNING":
		if quiet {
			return
		}
	}
}
