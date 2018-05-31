package repositories

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"condorbot/logger"
)

type MatchCase struct{
	MatchExact bool
	Request string
	Response string
}

type Repository interface{
	GetExactMatchMap() map[string]string
	GetWordMatchMap() map[string]string
}

type FireBaseRepository struct{
	Delegate func (string) (*http.Response, error)
	Logger logger.Logger
}

func (repo FireBaseRepository) GetExactMatchMap(url string) map[string]string {
	resp, err := repo.Delegate(url)
	if err != nil {
		repo.Logger.Err("FireBaseRepository", "First err: "+err.Error())
	}
	defer resp.Body.Close()

	var bytesArray []byte
	if resp.StatusCode == http.StatusOK {
		bytesArray, err = ioutil.ReadAll(resp.Body)
		if err != nil{
			repo.Logger.Err("FireBaseRepository", "First err: "+err.Error())
		}
	}

	if bytesArray != nil{
		var cases []MatchCase
		json.Unmarshal(bytesArray, &cases)
		return ExactMatchCasesToMap(cases)
	}

	return nil
}

func (repo FireBaseRepository) GetWordMatchMap(url string) map[string]string {
	resp, err := repo.Delegate(url)
	if err != nil {
		repo.Logger.Err("FireBaseRepository", "First err: "+err.Error())
	}
	defer resp.Body.Close()

	var bytesArray []byte
	if resp.StatusCode == http.StatusOK {
		bytesArray, err = ioutil.ReadAll(resp.Body)
		if err != nil{
			repo.Logger.Err("FireBaseRepository", "First err: "+err.Error())
		}
	}

	if bytesArray != nil{
		var cases []MatchCase
		json.Unmarshal(bytesArray, &cases)
		return WordMatchCasesToMap(cases)
	}

	return nil
}


func ExactMatchCasesToMap(matchCases []MatchCase) map[string]string{
	dict := make(map[string]string)
	for _,matchCase := range matchCases {
		if matchCase.MatchExact {
			dict[matchCase.Request] = matchCase.Response
		}
	}
	return dict
}

func WordMatchCasesToMap(matchCases []MatchCase) map[string]string{
	dict := make(map[string]string)
	for _,matchCase := range matchCases {
		if !matchCase.MatchExact {
			dict[matchCase.Request] = matchCase.Response
		}
	}
	return dict
}