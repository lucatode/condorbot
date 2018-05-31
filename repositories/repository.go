package repositories

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type Repository interface{
	GetExactMatchMap() map[string]string
}

type FireBaseRepository struct{
	Delegate func (string) (*http.Response, error)
}

func (repo FireBaseRepository) GetExactMatchMap(url string) map[string]string {
	resp, err := repo.Delegate(url)
	if err != nil {
		// err
	}
	defer resp.Body.Close()

	var bytesArray []byte
	if resp.StatusCode == http.StatusOK {
		bytesArray, err = ioutil.ReadAll(resp.Body)
		if err != nil{
			//err
		}
	}

	if bytesArray != nil{
		var cases []MatchCase
		json.Unmarshal(bytesArray, &cases)
		return MatchCasesToMap(cases)
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