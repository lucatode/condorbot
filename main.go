package main

import (
	"condorbot/initializer"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/telegram-bot-api.v4"
)

func Init() initializer.Initializer {
	return initializer.NewInitializer(initializer.NewEnvReader())
}

func NotifyChannel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	message := "Notifing channel: " + params["channel"]
	w.Write([]byte(message))
}

// func GetMapFromFirebase() map[string]string {
// 	client := &http.Client{}
// 	matchCases := repo.GetMatchCases(os.Getenv("FirebaseResponses"), client) //TODO: use initializer - "https://condorbot-c36af.firebaseio.com/responses.json" - GetServiceUrl()
// 	return repo.MapMatchCases(matchCases)
// }

func main() {
	//INIT
	initializer := Init()

	// SETUP BOT
	bot, err := tgbotapi.NewBotAPI(initializer.GetApiToken())
	if err != nil {
		log.Fatal(err)
	}

	// BOT CONFIG
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(initializer.GetServerUrl() + bot.Token))
	if err != nil {
		log.Fatal(err)
	}

	// SETUP INPUT ROUTES
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/notify/{channel}", NotifyChannel).Methods("GET")
	go http.ListenAndServe(":"+port, nil)

	// FETCH MESSAGES
	updates := bot.ListenForWebhook("/" + bot.Token)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// //  REGISTERING TO A CHANNEL

		// //  REPLY TO A MESSAGE
		if update.Message.Text == "Test" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			bot.Send(msg)
		}
	}
}
