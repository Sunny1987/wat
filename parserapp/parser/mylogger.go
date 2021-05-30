package parser

import (
	"log"
)

type MyLogger struct {
	l *log.Logger
}

func GetMyLogger(l *log.Logger) *MyLogger {
	return &MyLogger{l: l}
}
