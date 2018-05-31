package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func MockExactMatchDictionary() map[string]string {
	return map[string]string{
		"notify":"notified",
	}
	}

func MockWordMatcherDictionary() map[string]string {
	return map[string]string{
		"abba": "Found abba",
		"bccb": "Found bccb",
		"cddc": "Found cddc",
	}
}

//func MockCommandDictionary() map[string]string {
//	return map[string]string{
//		"#notify":"notified",
//		"#subscribe":"subscribed",
//	}
//}
//
//func TestCheckCommandMatch(t *testing.T) {
//	matchOutput, stringOutput := NewCommandsMatcher(MockCommandDictionary()).MatchString("#notify")
//	assert.Equal(t, stringOutput, "notified", "Getting specific string")
//	assert.Equal(t, matchOutput, true, "Getting specific match bool")
//}

func TestCheckStringMatch(t *testing.T) {
	matchOutput, stringOutput := NewExactMatcher(MockExactMatchDictionary()).MatchString("notify")
	assert.Equal(t, stringOutput, "notified", "Getting specific string")
	assert.Equal(t, matchOutput, true, "Getting specific match bool")
}

func TestContainsWordMatch(t *testing.T) {
	matchOutput, stringOutput := NewContainsWordMatcher(MockWordMatcherDictionary()).MatchString("abba abab ababa")
	assert.Equal(t,"Found abba", stringOutput, "Getting specific string")
	assert.Equal(t, matchOutput, true, "Getting specific match bool")

	matchOutput, stringOutput = NewContainsWordMatcher(MockWordMatcherDictionary()).MatchString("aaaa bccb")
	assert.Equal(t, "Found bccb",stringOutput,  "Getting specific string")
	assert.Equal(t, matchOutput, true, "Getting specific match bool")

	matchOutput, stringOutput = NewContainsWordMatcher(MockWordMatcherDictionary()).MatchString("cddc")
	assert.Equal(t,"", stringOutput,  "Getting specific string")
	assert.Equal(t, matchOutput, false, "Getting specific match bool")
}

func TestExactMatchDecorated(t *testing.T) {

	matcher := ContainsWordDecorated(
		MockWordMatcherDictionary(),
		NewExactMatcher(MockExactMatchDictionary()),
		)
	matchOutput, stringOutput := matcher.MatchString("aaaa bccb")
	assert.Equal(t, "Found bccb",stringOutput,  "Getting specific string")
	assert.Equal(t, matchOutput, true, "Getting specific match bool")

	matchOutput, stringOutput = ContainsWordDecorated(
		MockWordMatcherDictionary(),
		NewExactMatcher(MockExactMatchDictionary(),
			)).MatchString("notify")
	assert.Equal(t,"notified", stringOutput,  "Getting specific string")
	assert.Equal(t, matchOutput, true, "Getting specific match bool")
}


