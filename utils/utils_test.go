package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const TEST_POST_URL string = ""
const TEST_PUT_URL string = ""
const JSON_CONTENT_FOR_POST string = ""
const JSON_CONTENT_FOR_PUT string = ""

func TestSuite(t *testing.T) {
	assert.Nil(t, nil)
}

func TestJsonPostNotPanic(t *testing.T) {
	assert.NotPanics(t, func() { JsonPost(TEST_POST_URL, JSON_CONTENT_FOR_POST) })
}

func TestJsonPutNotPanic(t *testing.T) {
	assert.NotPanics(t, func() { JsonPut(TEST_PUT_URL, JSON_CONTENT_FOR_PUT) })
}