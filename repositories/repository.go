package repositories

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"condorbot/logger"
	"strconv"
)

type Repository interface{
	GetExactMatchMap() map[string]string
}

type FireBaseRepository struct{
	Delegate func (string) (*http.Response, error)
	Logger logger.Logger
}

func (repo FireBaseRepository) GetExactMatchMap(url string) map[string]string {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Custom-Header", "log")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		repo.Logger.Err("FireBaseRepository", "First err: "+err.Error())
	}
	defer resp.Body.Close()

	var bytesArray []byte
	if resp.StatusCode == http.StatusOK {
		bytesArray, err = ioutil.ReadAll(resp.Body)
		if err != nil{
			repo.Logger.Err("FireBaseRepository", "Second err: "+err.Error())
		}
	}

	if bytesArray != nil{
		var cases []MatchCase
		json.Unmarshal(bytesArray, &cases)
		for _, e := range cases {
			repo.Logger.Log("FireBaseRepository", "Loaded: "+e.Request)
		}

		dict := make(map[string]string)
		for _,matchCase := range cases {
			if matchCase.MatchExact {
				dict[matchCase.Request] = matchCase.Response
				repo.Logger.Log("FireBaseRepository", "To Map: "+matchCase.Request)
			}
		}

		l := len(dict)
		repo.Logger.Log("FireBaseRepository", "Dict created len: "+strconv.Itoa(l))
		return dict

	}

	return nil
}

type MatchCase struct{
	MatchExact bool
	Request string
	Response string
}

func MatchCasesToMap(matchCases []MatchCase) map[string]string{
	dict := make(map[string]string)
	for _,matchCase := range matchCases {
		if matchCase.MatchExact {
			dict[matchCase.Request] = matchCase.Response
		}
	}
	return dict
}