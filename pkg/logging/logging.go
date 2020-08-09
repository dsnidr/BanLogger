package logging

import "log"

// Info logs out an Info level message
func Info(source string, message string) {
	writeLog("INFO", source, message)
}

// Warn logs out an Warn level message
func Warn(source string, message string) {
	writeLog("WARN", source, message)
}

// Error logs out an Error level message
func Error(source string, message string) {
	writeLog("ERROR", source, message)
}

func writeLog(level string, source string, message string) {
	log.Printf("[%s] %s: %s\n", level, source, message)
}
