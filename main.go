package main

import (
	"condorbot/initializer"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/telegram-bot-api.v4"
)

func Init(args []string) initializer.Initializer {
	if len(args) > 1 {
		return initializer.NewInitializer(initializer.NewJsonReader(args[1]))
	}
	return initializer.NewInitializer(initializer.NewJsonReader("initializer_test.json"))
}

func NotifyChannel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	message := "Notifing channel: " + params["channel"]
	w.Write([]byte(message))
}

func main() {
	//INIT
	//initializer := Init(os.Args)

	// SETUP BOT
	bot, err := tgbotapi.NewBotAPI(os.Getenv("ApiToken"))
	if err != nil {
		log.Fatal(err)
	}

	// BOT CONFIG
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(os.Getenv("ServerUrl") + "" + bot.Token))
	if err != nil {
		log.Fatal(err)
	}

	// SETUP INPUT ROUTES
	port := os.Getenv("PORT")
	port = "8080"
	router := mux.NewRouter()
	router.HandleFunc("/notify/{channel}", NotifyChannel).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))

	// FETCH MESSAGES
	updates := bot.ListenForWebhook("/" + bot.Token)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// //  REGISTERING TO A CHANNEL

		// //  REPLY TO A MESSAGE
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		bot.Send(msg)
	}

}
