package main

import (
	"condor/initializer"
	"os"
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func Init(args []string) initializer.Initializer {
	if len(args) > 1 {
		return initializer.NewInitializer(initializer.NewJsonReader(args[1]))
	}
	return initializer.NewInitializer(initializer.NewJsonReader("initializer_test.json"))
}

func main(){
	//INIT
	initializer := Init(os.Args)

	// SETUP BOT
	bot, err := tgbotapi.NewBotAPI(initializer.GetApiToken())
	if err != nil {
		log.Fatal(err)
	}

	// BOT CONFIG
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(initializer.GetServerUrl() + "" + bot.Token))
	if err != nil {
	log.Fatal(err)
	}


	//port := os.Getenv("PORT")

}
