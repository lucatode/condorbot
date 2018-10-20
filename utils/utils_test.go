package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSuite(t *testing.T) {
	assert.Nil(t, nil)
}

func TestJsonPostNotPanic(t *testing.T){
	assert.NotPanics(t, func(){ JsonPost("",".....") })
}