package initializer

import (
	"os"
	"strconv"
)

func NewInitializer(ps InitializerStorage) *ParameterInitializer {
	return &ParameterInitializer{
		storage: ps,
	}
}

type Initializer interface {
	GetApiToken() string
	GetServerUrl() string
	GetTimerSeconds() int
	GetConfigResponsesUrl() string
	GetLoggerServiceUri() string
	GetConfigSubscriptionsUrl() string
}

type ParameterInitializer struct {
	storage InitializerStorage
}

func (init ParameterInitializer) GetLoggerServiceUri() string {
	return init.storage.GetData()["ConfigLogsUrl"]
}

func (init ParameterInitializer) GetApiToken() string {
	return init.storage.GetData()["ApiToken"]
}

func (init ParameterInitializer) GetServerUrl() string {
	return init.storage.GetData()["ServerUrl"]
}

func (init ParameterInitializer) GetTimerSeconds() int {
	ret, _ := strconv.Atoi(init.storage.GetData()["TimerSeconds"])
	return ret
}

func (init ParameterInitializer) GetConfigResponsesUrl() string {
	return init.storage.GetData()["ConfigResponsesUrl"]
}

func (init ParameterInitializer) GetConfigSubscriptionsUrl() string {
	return init.storage.GetData()["ConfigSubscriptionsUrl"]
}

//--- Initializer Storage
type InitializerStorage interface {
	GetData() map[string]string
}

//--- Env Storage
func NewEnvReader() *EnvStorage {
	return &EnvStorage{""}
}

type EnvStorage struct {
	storageType string
}

func (storage EnvStorage) GetData() map[string]string {
	return map[string]string{
		"ApiToken":                 os.Getenv("ApiToken"),
		"ServerUrl":                os.Getenv("ServerUrl"),
		"TimerSeconds":             os.Getenv("TimerSeconds"),
		"ConfigLogsUrl":         os.Getenv("ConfigLogsUrl"),
		"ConfigResponsesUrl":     	os.Getenv("ConfigResponsesUrl"),
		"ConfigSubscriptionsUrl": 	os.Getenv("ConfigSubscriptionsUrl"),
	}
}