package logger

import (
	"time"
)

type Log struct{
	Source string
	Message string
	Level string
	Time string
}

type Logger interface{
	Log(source string, message string, f func (url string, log interface{}))
	Warn(source string, message string, f func (url string, log interface{}))
	Err(source string, message string, f func (url string, log interface{}))
}

type FirebaseLogger struct{
	Url string
}

func (logger FirebaseLogger) Log(source string, message string, f func (url string, log interface{})) {
	l := Log { Source: source, Message: message, Level:"Info", Time:time.Now().String() }
	f(logger.Url, l)
}

func (logger FirebaseLogger) Warn(source string, message string, f func (url string, log interface{})) {
	l := Log { Source: source, Message: message, Level:"Warning", Time:time.Now().String() }
	f(logger.Url, l)
}

func (logger FirebaseLogger) Err(source string, message string, f func (url string, log interface{})) {
	l := Log { Source: source, Message: message, Level:"Error", Time:time.Now().String() }
	f(logger.Url, l)
}

type PutLogger struct{
	endPoint string
}

func (logger PutLogger) Log(source string, message string, f func (url string, log interface{})) {
	l := Log { Source: source, Message: message, Level:"Info", Time:time.Now().String() }
	f(logger.endPoint, l)
}

func (logger PutLogger) Warn(source string, message string, f func (url string, log interface{})) {
	l := Log { Source: source, Message: message, Level:"Warning", Time:time.Now().String() }
	f(logger.endPoint, l)
}

func (logger PutLogger) Err(source string, message string, f func (url string, log interface{})) {
	l := Log { Source: source, Message: message, Level:"Error", Time:time.Now().String() }
	f(logger.endPoint, l)
}