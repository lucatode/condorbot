package main

import (
	"condorbot/dispatcher"
	"condorbot/initializer"
	"condorbot/subscriber"
	"log"
	"net/http"
	"os"
	"strconv"

	"condorbot/logger"
	"condorbot/parser"
	"condorbot/repositories"
	"condorbot/utils"
	"encoding/json"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"strings"
)

func Init() initializer.Initializer {
	return initializer.NewInitializer(initializer.NewEnvReader())
}
func CreateLogger(init initializer.Initializer) logger.PutLogger {
	logger := logger.PutLogger{init.GetLoggerServiceUri(), utils.JsonPost}
	logger.Log("MAIN", "Starting")
	return logger
}
func CreateRepository(logger logger.PutLogger) repositories.ConfigRepository {
	client := http.Client{}
	return repositories.ConfigRepository{client.Get, logger}
}
func BuildMessage(message *tgbotapi.Message) parser.Message {
	return parser.Message{message.Text, strconv.FormatInt(message.Chat.ID, 10)}
}
func BuildCommandDispatcher(url string) dispatcher.Dispatcher {
	return dispatcher.CommandDispatcher{map[string]func([]string, string) string{
		"#subscribe": func(split []string, chatId string) string { return subscriber.AddSubscription(url, split, chatId, utils.JsonPost) },
	}}
}
func NotifyHandler(init initializer.Initializer, bot *tgbotapi.BotAPI) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		channel := strings.TrimPrefix(r.URL.Path, "/notify/")
		channelsToNotify := subscriber.GetChatIdsForChannel(init.GetConfigSubscriptionsUrl(), channel)

		type NotificationMessage struct {
			Key     string
			Source  string
			Message string
		}

		var mex NotificationMessage

		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Error reading request body",
					http.StatusInternalServerError)
			}

			if body != nil {
				json.Unmarshal(body, &mex)
			}
		}

		if mex.Key == os.Getenv("SecuriyKey") {
			for _, c := range channelsToNotify {
				i, _ := strconv.ParseInt(c, 10, 64)
				msg := tgbotapi.NewMessage(i, "["+mex.Source+"]: "+mex.Message)
				bot.Send(msg)
			}
		}
	}
}

func main() {
	//INIT
	init := Init()
	logger := CreateLogger(init)
	repo := CreateRepository(logger)
	m := repo.GetWordMatchMap(init.GetConfigResponsesUrl())
	p := parser.CommandsDecorated(
		BuildCommandDispatcher(init.GetConfigSubscriptionsUrl()),
		parser.ContainsWordDecorated(m,
			parser.NewExactMatcher(
				repo.GetExactMatchMap(init.GetConfigResponsesUrl()))))

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
	http.HandleFunc("/notify/", NotifyHandler(init, bot))

	// FETCH MESSAGES
	updates := bot.ListenForWebhook("/" + bot.Token)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		ok, text := p.ParseMessage(BuildMessage(update.Message))
		if ok {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			bot.Send(msg)
		}
	}
}
