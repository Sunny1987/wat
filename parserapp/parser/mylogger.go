package parser

import (
	"log"
)

//MyLogger is the parser package logger
type MyLogger struct {
	l      *log.Logger
	req    interface{}
	person string
}

//GetMyLogger is the parser package logger function
func GetMyLogger(l *log.Logger, req interface{}, person string) *MyLogger {
	return &MyLogger{l: l, req: req, person: person}
}
