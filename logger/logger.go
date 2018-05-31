package logger

import "github.com/cosn/firebase"

type Log struct{
	Source string
	Message string
	Level string
}

type Logger interface{
	Log(source string, message string)
	Warn(source string, message string)
	Err(source string, message string)
}

type FirebaseLogger struct{
	Url string
}

func (l FirebaseLogger) Log(source string, message string) {
	firebase := new(firebase.Client)
	firebase.Init(l.Url, "", nil)

	n := &Log { Source: "Source", Message: "Message", Level:"Info" }
	firebase.Push(n, nil)
}

func (l FirebaseLogger) Warn(source string, message string) {
	firebase := new(firebase.Client)
	firebase.Init(l.Url, "", nil)

	n := &Log { Source: "Source", Message: "Message", Level:"Warning" }
	firebase.Push(n, nil)
}

func (l FirebaseLogger) Err(source string, message string) {
	firebase := new(firebase.Client)
	firebase.Init(l.Url, "", nil)

	n := &Log { Source: "Source", Message: "Message", Level:"Error" }
	firebase.Push(n, nil)
}

