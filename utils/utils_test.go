package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuite(t *testing.T) {
	assert.Nil(t, nil)
}

func TestJsonPostNotPanic(t *testing.T) {
	assert.NotPanics(t, func() { JsonPost("", ".....") })
}
