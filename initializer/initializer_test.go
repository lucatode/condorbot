package initializer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAPITocken(t *testing.T) {
	i := NewInitializer(NewJsonReader("initializer.json"))

	token := i.GetApiToken()

	assert.Equal(t, "ABCD012345678", token, "Get Api Token from mocked storage")
}

func TestGetServerUrl(t *testing.T) {
	i := NewInitializer(NewJsonReader("initializer.json"))

	url := i.GetServerUrl()

	assert.Equal(t, "https://xxxxx.yyyyyy.com:443/", url, "Get URL from mocked storage")
}

func TestGetTimerSeconds(t *testing.T) {
	i := NewInitializer(NewJsonReader("initializer.json"))

	seconds := i.GetTimerSeconds()

	assert.Equal(t, 3600, seconds, "Get URL from mocked storage")
}

func TestReadDifferentValuesFromDifferentFiles(t *testing.T) {
	it1 := NewInitializer(NewJsonReader("initializer_test_1.json"))
	it2 := NewInitializer(NewJsonReader("initializer_test_2.json"))

	seconds1 := it1.GetTimerSeconds()
	seconds2 := it2.GetTimerSeconds()

	assert.Equal(t, 3600, seconds1, "Get URL from mocked storage")
	assert.Equal(t, 100, seconds2, "Get URL from mocked storage")
}
