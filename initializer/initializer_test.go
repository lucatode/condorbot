package initializer

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func EnvVarInitializer() {
	os.Setenv("ApiToken", "ABCD012345678")
	os.Setenv("ServerUrl", "https://xxxxx.yyyyyy.com:443/")
	os.Setenv("TimerSeconds", "3600")
	os.Setenv("FirebaseResponsesUrl", "https://xxxxx.yyyyyy.com:443/")
}

func TestGetAPITocken(t *testing.T) {

	EnvVarInitializer()

	i := NewInitializer(NewEnvReader())

	token := i.GetApiToken()

	assert.Equal(t, "ABCD012345678", token, "Get Api Token from mocked storage")
}

func TestGetServerUrl(t *testing.T) {
	EnvVarInitializer()

	i := NewInitializer(NewEnvReader())

	url := i.GetServerUrl()

	assert.Equal(t, "https://xxxxx.yyyyyy.com:443/", url, "Get URL from mocked storage")
}

func TestGetTimerSeconds(t *testing.T) {
	EnvVarInitializer()

	i := NewInitializer(NewEnvReader())

	seconds := i.GetTimerSeconds()

	assert.Equal(t, 3600, seconds, "Get URL from mocked storage")
}
