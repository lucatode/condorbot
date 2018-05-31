package main

import (
	"condorbot/initializer"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/telegram-bot-api.v4"
	"condorbot/parser"
	"condorbot/repositories"
	"condorbot/logger"
)

func Init() initializer.Initializer {
	return initializer.NewInitializer(initializer.NewEnvReader())
}

func NotifyChannel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	message := "Notifing channel: " + params["channel"]
	w.Write([]byte(message))
}

func main() {
	//INIT
	init := Init()
	client := http.Client{}
	getFunc := client.Get
	logger := logger.FirebaseLogger{init.GetFireBaseLogsUrl()}
	logger.Log("MAIN", "Starting")
	repo := repositories.FireBaseRepository{getFunc, logger}
	parser := parser.NewExactMatcher(repo.GetExactMatchMap(init.GetFireBaseResponsesUrl()))

	// SETUP BOT
	bot, err := tgbotapi.NewBotAPI(init.GetApiToken())
	if err != nil {
		log.Fatal(err)
	}

	// BOT CONFIG
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(init.GetServerUrl() + bot.Token))
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

		ok,text := parser.MatchString(update.Message.Text)

		if ok {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			bot.Send(msg)
		}
	}
}
