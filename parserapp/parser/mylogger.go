package parser

import (
	"log"
)

//MyLogger is the parser package logger
type MyLogger struct {
	l *log.Logger
}

//GetMyLogger is the parser package logger function
func GetMyLogger(l *log.Logger) *MyLogger {
	return &MyLogger{l: l}
}
