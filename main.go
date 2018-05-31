package main

import (
	"condorbot/initializer"
	"log"
	"net/http"
	"os"

	"gopkg.in/telegram-bot-api.v4"
	"condorbot/parser"
	"condorbot/repositories"
	"condorbot/logger"
)

func Init() initializer.Initializer {
	return initializer.NewInitializer(initializer.NewEnvReader())
}

func NotifyChannel(w http.ResponseWriter, r *http.Request) {
	message := "Notifing channel: "
	w.Write([]byte(message))
}

func CreateLogger(init initializer.Initializer) logger.FirebaseLogger {
	logger := logger.FirebaseLogger{init.GetFireBaseLogsUrl()}
	logger.Log("MAIN", "Starting")
	return logger
}

func CreateRepository(logger logger.FirebaseLogger) repositories.FireBaseRepository {
	client := http.Client{}
	return repositories.FireBaseRepository{client.Get, logger}
}

func main() {
	//INIT
	init := Init()
	logger := CreateLogger(init)
	repo := CreateRepository(logger)
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
	go http.ListenAndServe(":"+port, nil)

	http.HandleFunc("/notify/{channel}", func(w http.ResponseWriter, r *http.Request) {
		NotifyChannel(w,r)
	})

	// FETCH MESSAGES
	updates := bot.ListenForWebhook("/" + bot.Token)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		ok, text := parser.MatchString(update.Message.Text)

		if ok {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			bot.Send(msg)
		}
	}
}


