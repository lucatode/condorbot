package parser

import (
	"condorbot/dispacher"
	"strings"
)

type Parser interface {
	ParseMessage(Message) (bool, string)
}

type Message struct {
	Text   string
	ChatId string
}

type ExactMatcher struct {
	exactMatchDict map[string]string
}

func (em ExactMatcher) ParseMessage(message Message) (bool, string) {
	val, ok := em.exactMatchDict[message.Text]
	return ok, val
}

func NewExactMatcher(dict map[string]string) Parser {
	return ExactMatcher{dict}
}

type ContainsWordMatcher struct {
	delegate         func(Message) (bool, string)
	containsWordDict map[string]string
}

func (cwm ContainsWordMatcher) ParseMessage(message Message) (bool, string) {
	inputString := message.Text
	if strings.Contains(inputString, " ") {
		splittedMessage := strings.Split(inputString, " ")
		for _, word := range splittedMessage {
			val, ok := cwm.containsWordDict[word]
			if ok {
				return ok, val
			}
		}
	}
	return cwm.delegate(message)
}

func NewContainsWordMatcher(dict map[string]string) Parser {
	delegate := func(input Message) (bool, string) { return false, "" }
	return ContainsWordMatcher{delegate, dict}
}

func ContainsWordDecorated(dict map[string]string, matcher Parser) Parser {
	return ContainsWordMatcher{matcher.ParseMessage, dict}
}

type CommandsMatcher struct {
	delegate  func(Message) (bool, string)
	dispacher dispacher.Dispacher
}

func (cm CommandsMatcher) ParseMessage(message Message) (bool, string) {
	inputString := message.Text
	splittedMessage := strings.Split(inputString, " ")
	ok, f := cm.dispacher.GetActionFunc(splittedMessage[0])
	if ok {

		return ok, f(splittedMessage, message.ChatId)
	}

	return cm.delegate(message)
}

func NewCommandsMatcher(dispacher dispacher.Dispacher) Parser {
	delegate := func(input Message) (bool, string) { return false, "" }
	return CommandsMatcher{delegate, dispacher}
}

func CommandsDecorated(dispacher dispacher.Dispacher, matcher Parser) Parser {
	return CommandsMatcher{matcher.ParseMessage, dispacher}
}
