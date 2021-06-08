package parser

import (
	"log"
)

//MyLogger is the parser package logger
type MyLogger struct {
	l   *log.Logger
	req interface{}
}

//GetMyLogger is the parser package logger function
func GetMyLogger(l *log.Logger, req interface{}) *MyLogger {
	return &MyLogger{l: l, req: req}
}
