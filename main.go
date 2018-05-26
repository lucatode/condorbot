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
	_, err := tgbotapi.NewBotAPI(initializer.GetApiToken())
	if err != nil {
		log.Fatal(err)
	}

}
