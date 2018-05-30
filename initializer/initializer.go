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
}

type ParameterInitializer struct {
	storage InitializerStorage
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
		"ApiToken":             os.Getenv("ApiToken"),
		"ServerUrl":            os.Getenv("ServerUrl"),
		"TimerSeconds":         os.Getenv("TimerSeconds"),
		"FirebaseResponsesUrl": os.Getenv("FirebaseResponsesUrl"),
	}
}
