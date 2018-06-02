package subscriber

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAPITocken(t *testing.T) {
	message := "#subscribe channel xxxx"
	split := strings.Split(message, " ")
	response := AddSubscription("https://condorbot-c36af.firebaseio.com/subscriptions", split, "chatId 1")
	AddSubscription("https://condorbot-c36af.firebaseio.com/subscriptions", split, "chatId 2")
	AddSubscription("https://condorbot-c36af.firebaseio.com/subscriptions", split, "chatId 3")

	assert.Equal(t, "Subscribed this chat to channel xxxx", response, "Getting specific match bool")
}
