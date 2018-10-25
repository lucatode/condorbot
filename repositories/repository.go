package repositories

import (
	"condorbot/logger"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type MatchCase struct {
	MatchExact bool
	Request    string
	Response   string
}

type Repository interface {
	GetExactMatchMap() map[string]string
	GetWordMatchMap() map[string]string
}

type ConfigRepository struct {
	Delegate func(string) (*http.Response, error)
	logger   logger.Logger
}

func (repo ConfigRepository) GetExactMatchMap(url string) map[string]string {
	resp, err := repo.Delegate(url)
	if err != nil {
		repo.logger.Err("ConfigRepository", "First err: "+err.Error())
	}
	defer resp.Body.Close()

	var byteArray []byte
	if resp.StatusCode == http.StatusOK {
		byteArray, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			repo.logger.Err("ConfigRepository", "First err: "+err.Error())
		}
	}

	if byteArray != nil {
		var cases []MatchCase
		s := string(byteArray)
		repo.logger.Log("REPO", "received "+s)
		json.Unmarshal(byteArray, &cases)
		return exactMatchCasesToMap(cases)
	}

	return nil
}

func (repo ConfigRepository) GetWordMatchMap(url string) map[string]string {
	resp, err := repo.Delegate(url)
	if err != nil {
		repo.logger.Err("ConfigRepository", "First err: "+err.Error())
	}
	defer resp.Body.Close()

	var bytesArray []byte
	if resp.StatusCode == http.StatusOK {
		bytesArray, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			repo.logger.Err("ConfigRepository", "First err: "+err.Error())
		}
	}

	if bytesArray != nil {
		var cases []MatchCase
		json.Unmarshal(bytesArray, &cases)
		return wordMatchCasesToMap(cases)
	}

	return nil
}

func exactMatchCasesToMap(matchCases []MatchCase) map[string]string {
	dict := make(map[string]string)
	for _, matchCase := range matchCases {
		if matchCase.MatchExact {
			dict[matchCase.Request] = matchCase.Response
		}
	}
	return dict
}

func wordMatchCasesToMap(matchCases []MatchCase) map[string]string {
	dict := make(map[string]string)
	for _, matchCase := range matchCases {
		if !matchCase.MatchExact {
			dict[matchCase.Request] = matchCase.Response
		}
	}
	return dict
}
