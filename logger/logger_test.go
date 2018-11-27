package logger

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPutLogger_Log(t *testing.T) {
	logger := BuildLogger()
	logger.Log("TEST_SUITE", "This is a log message")
}

func TestPutLogger_Warn(t *testing.T) {
	logger := BuildLogger()
	logger.Warn("TEST_SUITE", "This is a log message")
}

func TestPutLogger_Err(t *testing.T) {
	logger := BuildLogger()
	logger.Err("TEST_SUITE", "This is a log message")
}


func BuildLogger() Logger{
	return PutLogger{"", FAKE_PUT}
}

func FAKE_PUT(url string, log interface{}){

}