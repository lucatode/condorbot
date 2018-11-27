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
	os.Setenv("ConfigResponsesUrl", "https://xxxxx.Config.com:443/")
	os.Setenv("ConfigLogsUrl", "https://xxxxx.Config.com:443/logs")
	os.Setenv("ConfigSubscriptionsUrl", "https://xxxxx.Config.com:443/subscriptions")
}

func TestGetAPIToken(t *testing.T) {

	EnvVarInitializer()

	i := NewInitializer(NewEnvReader())

	token := i.GetApiToken()

	assert.Equal(t, "ABCD012345678", token, "Get Api Token from mocked storage")
}

func TestGetServerUrl(t *testing.T) {
	EnvVarInitializer()

	i := NewInitializer(NewEnvReader())

	url := i.GetServerUrl()

	assert.Equal(t, "https://xxxxx.yyyyyy.com:443/", url, "")
}

func TestGetConfigResponsesUrl(t *testing.T) {
	EnvVarInitializer()

	i := NewInitializer(NewEnvReader())

	url := i.GetConfigResponsesUrl()

	assert.Equal(t, "https://xxxxx.Config.com:443/", url, "")
}

func TestGetConfigLogsUrl(t *testing.T) {
	EnvVarInitializer()

	i := NewInitializer(NewEnvReader())

	url := i.GetLoggerServiceUri()

	assert.Equal(t, "https://xxxxx.Config.com:443/logs", url, "")
}

func TestGetConfigSubscriptionsUrl(t *testing.T) {
	EnvVarInitializer()

	i := NewInitializer(NewEnvReader())

	url := i.GetConfigSubscriptionsUrl()

	assert.Equal(t, "https://xxxxx.Config.com:443/subscriptions", url, "")
}

func TestGetTimerSeconds(t *testing.T) {
	EnvVarInitializer()

	i := NewInitializer(NewEnvReader())

	seconds := i.GetTimerSeconds()

	assert.Equal(t, 3600, seconds, "")
}
