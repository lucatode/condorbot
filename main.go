package main

import (
	"condorbot/dispacher"
	"condorbot/initializer"
	"log"
	"net/http"
	"os"

	"condorbot/logger"
	"condorbot/parser"
	"condorbot/repositories"
	"strings"

	"gopkg.in/telegram-bot-api.v4"
)

func Init() initializer.Initializer {
	return initializer.NewInitializer(initializer.NewEnvReader())
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

func BuildCommandDispacher() dispacher.Dispacher {
	return dispacher.CommandDispacher{map[string]func([]string) string{
		"#register": func(split []string) string { return "" },
	}}
}

func main() {
	//INIT
	init := Init()
	logger := CreateLogger(init)
	repo := CreateRepository(logger)

	parser := parser.CommandsDecorated(
		BuildCommandDispacher(),
		parser.ContainsWordDecorated(
			repo.GetWordMatchMap(init.GetFireBaseResponsesUrl()),
			parser.NewExactMatcher(
				repo.GetExactMatchMap(init.GetFireBaseResponsesUrl()))))

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
	logger.Log("MAIN", "port: "+port)
	go http.ListenAndServe(":"+port, nil)

	http.HandleFunc("/notify/", func(w http.ResponseWriter, r *http.Request) {
		channel := strings.TrimPrefix(r.URL.Path, "/notify/")
		logger.Log("MAIN", "call notify chan: "+channel)
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
