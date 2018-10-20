package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func JsonPost(url string, object interface{}) {
	jsonStr, _ := json.Marshal(object)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("X-Custom-Header", "object")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
