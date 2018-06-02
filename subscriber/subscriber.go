package subscriber

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Subscription struct {
	channel string
	chatId  string
}

func httpPost(url string, subscription Subscription) {
	jsonStr, _ := json.Marshal(subscription)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "log")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func AddSubscription(url string, message []string, chatId string) string {

	if len(message) == 3 {
		channel := message[2]
		s := Subscription{channel, chatId}
		httpPost(url, s)
		return "Subscribed this chat to channel " + channel
	}
	return "invalid message for subscribe channel"
}
