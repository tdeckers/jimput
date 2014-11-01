package log

import (
	"fmt"
	"os"
	"time"
)

var (
	tag     string
	quiet   = false
	verbose = false
	debug   = false
)

func init() {
	tag = os.Args[0]
}

// SetTag sets the tag used for logging.
func SetTag(t string) {
	tag = t
}

// SetQuiet sets quiet mode.
func SetQuiet(enable bool) {
	quiet = enable
}

// SetDebug sets debug mode
func SetDebug(enable bool) {
	debug = enable
}

// SetVerbose sets verbose mode.
func SetVerbose(enable bool) {
	verbose = enable
}

// Debug logs a message with severity DEBUG.
func Debug(msg string) {
	if debug {
		log("DEBUG", msg)
	}
}

// Error logs a message with severity ERROR.
func Error(msg string) {
	log("ERROR", msg)
}

// Fatal logs a message with severity ERROR and exists by a call to os.Exit().
func Fatal(msg string) {
	Error(msg)
	os.Exit(1)
}

// Info logs a message with severity INFO.
func Info(msg string) {
	log("INFO", msg)
}

// Notice logs a message with severity NOTICE.
func Notice(msg string) {
	if verbose || debug {
		log("NOTICE", msg)
	}
}

// Warning logs a message with severity WARNING.
func Warning(msg string) {
	log("WARNING", msg)
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
		w = os.Stdout
	case "ERROR":
		w = os.Stderr
	}
	fmt.Fprintf(w, "%s %s %s[%d]: %s %s\n",
		timestamp, hostname, tag, os.Getpid(), level, msg)
}
