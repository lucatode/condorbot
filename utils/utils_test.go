package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSuite(t *testing.T) {
	assert.Nil(t, nil)
}

func TestNotPanic(t *testing.T){
	assert.NotPanics(t, func(){ JsonPost("",".....") })
}

func TestMarshallFailure(t *testing.T){

}

func TestPostFailure(t *testing.T){

}