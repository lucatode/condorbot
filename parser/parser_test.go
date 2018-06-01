package parser

import (
	"condorbot/dispacher"
	"testing"

	"github.com/stretchr/testify/assert"
)

func MockExactMatchDictionary() map[string]string {
	return map[string]string{
		"notify": "notified",
	}
}

func MockWordMatcherDictionary() map[string]string {
	return map[string]string{
		"abba": "Found abba",
		"bccb": "Found bccb",
		"cddc": "Found cddc",
	}
}

func MockDispacher() dispacher.Dispacher {
	return dispacher.NewCommandDispacher(map[string]func([]string) string{
		"#command": func([]string) string { return "command received" },
		"#subscribe_2": func(params []string) string {
			msg := "subscribed"
			for i, p := range params {
				if i > 0 {
					msg = msg + " " + p
				}
			}
			return msg
		}})
}

func TestExactMatchDecoratedWithCommand(t *testing.T) {

	matcher := CommandsDecorated(
		MockDispacher(),
		NewExactMatcher(MockExactMatchDictionary()),
	)
	matchOutput, stringOutput := matcher.MatchString("notify")
	assert.Equal(t, "notified", stringOutput, "")
	assert.Equal(t, matchOutput, true, "")

	matchOutput, stringOutput = matcher.MatchString("#command")
	assert.Equal(t, "command received", stringOutput, "")
	assert.Equal(t, matchOutput, true, "")
}

func TestCheckCommandMatch(t *testing.T) {
	matchOutput, stringOutput := NewCommandsMatcher(MockDispacher()).MatchString("#command")
	assert.Equal(t, stringOutput, "command received", "")
	assert.Equal(t, matchOutput, true, "")
}

func TestCommandParameters(t *testing.T) {
	matchOutput, stringOutput := NewCommandsMatcher(MockDispacher()).MatchString("#subscribe_2 channel xxxx")
	assert.Equal(t, stringOutput, "subscribed channel xxxx", "")
	assert.Equal(t, matchOutput, true, "")
}

func TestCheckStringMatch(t *testing.T) {
	matchOutput, stringOutput := NewExactMatcher(MockExactMatchDictionary()).MatchString("notify")
	assert.Equal(t, stringOutput, "notified", "")
	assert.Equal(t, matchOutput, true, "")
}

func TestContainsWordMatch(t *testing.T) {
	matchOutput, stringOutput := NewContainsWordMatcher(MockWordMatcherDictionary()).MatchString("abba abab ababa")
	assert.Equal(t, "Found abba", stringOutput, "")
	assert.Equal(t, matchOutput, true, "")

	matchOutput, stringOutput = NewContainsWordMatcher(MockWordMatcherDictionary()).MatchString("aaaa bccb")
	assert.Equal(t, "Found bccb", stringOutput, "")
	assert.Equal(t, matchOutput, true, "")

	matchOutput, stringOutput = NewContainsWordMatcher(MockWordMatcherDictionary()).MatchString("cddc")
	assert.Equal(t, "", stringOutput, "")
	assert.Equal(t, matchOutput, false, "")
}

func TestExactMatchDecorated(t *testing.T) {

	matcher := ContainsWordDecorated(
		MockWordMatcherDictionary(),
		NewExactMatcher(MockExactMatchDictionary()),
	)
	matchOutput, stringOutput := matcher.MatchString("aaaa bccb")
	assert.Equal(t, "Found bccb", stringOutput, "")
	assert.Equal(t, matchOutput, true, "")

	matchOutput, stringOutput = ContainsWordDecorated(
		MockWordMatcherDictionary(),
		NewExactMatcher(MockExactMatchDictionary())).MatchString("notify")
	assert.Equal(t, "notified", stringOutput, "")
	assert.Equal(t, matchOutput, true, "")
}
