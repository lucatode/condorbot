package parser

import (
	"condorbot/dispacher"
	"strings"
)

type Parser interface {
	MatchString(string) (bool, string)
}

type ExactMatcher struct {
	exactMatchDict map[string]string
}

func (em ExactMatcher) MatchString(inputString string) (bool, string) {
	val, ok := em.exactMatchDict[inputString]
	return ok, val
}

func NewExactMatcher(dict map[string]string) Parser {
	return ExactMatcher{dict}
}

type ContainsWordMatcher struct {
	delegate         func(string) (bool, string)
	containsWordDict map[string]string
}

func (cwm ContainsWordMatcher) MatchString(inputString string) (bool, string) {
	if strings.Contains(inputString, " ") {
		splittedMessage := strings.Split(inputString, " ")
		for _, word := range splittedMessage {
			val, ok := cwm.containsWordDict[word]
			if ok {
				return ok, val
			}
		}
	}
	return cwm.delegate(inputString)
}

func NewContainsWordMatcher(dict map[string]string) Parser {
	delegate := func(input string) (bool, string) { return false, "" }
	return ContainsWordMatcher{delegate, dict}
}

func ContainsWordDecorated(dict map[string]string, matcher Parser) Parser {
	return ContainsWordMatcher{matcher.MatchString, dict}
}

type CommandsMatcher struct {
	delegate  func(string) (bool, string)
	dispacher dispacher.Dispacher
}

func (cm CommandsMatcher) MatchString(inputString string) (bool, string) {

	splittedMessage := strings.Split(inputString, " ")
	ok, f := cm.dispacher.GetActionFunc(splittedMessage[0])
	if ok {

		return ok, f(splittedMessage)
	}

	return cm.delegate(inputString)
}

func NewCommandsMatcher(dispacher dispacher.Dispacher) Parser {
	delegate := func(input string) (bool, string) { return false, "" }
	return CommandsMatcher{delegate, dispacher}
}

func CommandsDecorated(dispacher dispacher.Dispacher, matcher Parser) Parser {
	return CommandsMatcher{matcher.MatchString, dispacher}
}
