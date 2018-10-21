package logger

import (
	"testing"
	"github.com/stretchr/testify/assert"
)



func TestSuite(t *testing.T) {
	assert.Nil(t, nil)
}

func TestPutLogger_Log(t *testing.T) {
	logger := BuildLogger()
	logger.Log("TEST_SUITE", "This is a log message")
}

func TestPutLogger_Warn(t *testing.T) {

}

func TestPutLogger_Err(t *testing.T) {

}



func BuildLogger() Logger{
	return PutLogger{"", FAKE_PUT}
}

func FAKE_PUT(url string, log interface{}){

}