package subscriber

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Subscription struct {
	Channel string
	ChatId  string
}

func AddSubscription(url string, message []string, chatId string, f func(url string, subscription Subscription) ) string {

	if len(message) == 3 {
		channel := message[2]
		s := Subscription{channel, chatId}
		f(url+"/"+channel+".json", s)
		return "Subscribed this chat to channel " + channel
	}
	return "invalid message for subscribe channel"
}

func GetChatIdsForChannel(url string, channel string) []string {
	chatIds := []string{}
	client := http.Client{}
	resp, err := client.Get(url + "/" + channel + ".json")

	if err != nil {
		//repo.Logger.Err("FireBaseRepository", "First err: "+err.Error())
	}
	defer resp.Body.Close()

	var bytesArray []byte
	if resp.StatusCode == http.StatusOK {
		bytesArray, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			//repo.Logger.Err("FireBaseRepository", "First err: "+err.Error())
		}
	}

	if bytesArray != nil {
		var subscriptions map[string]Subscription
		json.Unmarshal(bytesArray, &subscriptions)

		for _, sub := range subscriptions {
			chatIds = append(chatIds, sub.ChatId)
		}
	}
	return chatIds
}
