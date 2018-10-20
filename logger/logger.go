package logger

import (
	"time"
)

type Log struct {
	Source  string
	Message string
	Level   string
	Time    string
}

type Logger interface {
	Log(source string, message string )
	Warn(source string, message string)
	Err(source string, message string)
}

type FirebaseLogger struct {
	url string
	f   func(url string, log interface{})
}

func (logger FirebaseLogger) Log(source string, message string) {
	l := Log{Source: source, Message: message, Level: "Info", Time: time.Now().String()}
	logger.f(logger.url, l)
}

func (logger FirebaseLogger) Warn(source string, message string) {
	l := Log{Source: source, Message: message, Level: "Warning", Time: time.Now().String()}
	logger.f(logger.url, l)
}

func (logger FirebaseLogger) Err(source string, message string) {
	l := Log{Source: source, Message: message, Level: "Error", Time: time.Now().String()}
	logger.f(logger.url, l)
}

type PutLogger struct {
	endPoint string
	f        func(url string, log interface{})
}

func (logger PutLogger) Log(source string, message string) {
	l := Log{Source: source, Message: message, Level: "Info", Time: time.Now().String()}
	logger.f(logger.endPoint, l)
}

func (logger PutLogger) Warn(source string, message string) {
	l := Log{Source: source, Message: message, Level: "Warning", Time: time.Now().String()}
	logger.f(logger.endPoint, l)
}

func (logger PutLogger) Err(source string, message string) {
	l := Log{Source: source, Message: message, Level: "Error", Time: time.Now().String()}
	logger.f(logger.endPoint, l)
}
