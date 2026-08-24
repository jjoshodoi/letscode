package logging

import (
	"log"
	"strings"
)

type Level int

const (
	Debug Level = iota
	Info
	Warn
	Error
)

var level = Info

// Init configures the logger level. Accepts: debug, info, warn, error.
func Init(levelStr string) {
	switch strings.ToLower(levelStr) {
	case "debug":
		level = Debug
	case "info", "":
		level = Info
	case "warn", "warning":
		level = Warn
	case "error":
		level = Error
	default:
		level = Info
	}
	log.SetFlags(log.LstdFlags | log.Lmsgprefix)
}

func Debugf(format string, v ...interface{}) {
	if level <= Debug {
		log.Printf("[DEBUG] "+format, v...)
	}
}

func Infof(format string, v ...interface{}) {
	if level <= Info {
		log.Printf("[INFO] "+format, v...)
	}
}

func Warnf(format string, v ...interface{}) {
	if level <= Warn {
		log.Printf("[WARN] "+format, v...)
	}
}

func Errorf(format string, v ...interface{}) {
	if level <= Error {
		log.Printf("[ERROR] "+format, v...)
	}
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf("[FATAL] "+format, v...)
}
