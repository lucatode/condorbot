package initializer

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
)

func NewInitializer(ps InitializerStorage) *FileBasedInitializer {
	return &FileBasedInitializer{
		storage: ps,
	}
}

type Initializer interface {
	GetApiToken() string
	GetServerUrl() string
	GetTimerSeconds() int
}

type FileBasedInitializer struct {
	storage InitializerStorage
}

func (init FileBasedInitializer) GetApiToken() string {
	return init.storage.GetData()["ApiToken"]
}

func (init *FileBasedInitializer) GetServerUrl() string {
	return init.storage.GetData()["ServerUrl"]
}

func (init *FileBasedInitializer) GetTimerSeconds() int {
	ret, _ := strconv.Atoi(init.storage.GetData()["TimerSeconds"])
	return ret
}

//--- Initializer Storage
type InitializerStorage interface {
	GetData() map[string]string
}

//--- Json Storage
func NewJsonReader(path string) *JsonStorage {
	return &JsonStorage{"", path}
}

type JsonStorage struct {
	storageType string
	path        string
}

func (storage JsonStorage) GetData() map[string]string {
	content, err := ioutil.ReadFile(storage.path)
	if err != nil {
		log.Fatal(err)
	}

	jsonMap := make(map[string]string)
	err = json.Unmarshal([]byte(content), &jsonMap)
	if err != nil {
		panic(err)
	}

	return jsonMap
}

//--- Mocked Storage
func NewMockedReader() *MockedStorage {
	return &MockedStorage{""}
}

type MockedStorage struct {
	storageType string
}

func (storage MockedStorage) GetData() map[string]string {
	return map[string]string{
		"ApiToken":     "ABCD012345678",
		"ServerUrl":    "https://xxxxx.yyyyyy.com:443/",
		"TimerSeconds": "3600",
	}
}
