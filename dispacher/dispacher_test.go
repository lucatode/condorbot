package dispacher

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewMockedCommandDispacher() Dispacher {
	return CommandDispacher{map[string]func([]string) string{
		"#subscribe": func([]string) string { return "subscribed" },
		"#subscribe_2": func(params []string) string {
			msg := "subscribed"
			for i, p := range params {
				if i > 0 {
					msg = msg + " " + p
				}
			}
			return msg
		},
	}}
}

func TestCommandMatched(t *testing.T) {
	d := NewMockedCommandDispacher()
	ok, _ := d.GetActionFunc("#subscribe_2")
	assert.Equal(t, true, ok, "Getting specific match bool")
}

func TestGetActionFunc(t *testing.T) {
	d := NewMockedCommandDispacher()
	_, f := d.GetActionFunc("#subscribe_2")
	input := "#subscribe_2 channel xxxx"
	splittedInput := strings.Split(input, " ")
	m := f(splittedInput)
	assert.Equal(t, "subscribed channel xxxx", m, "Getting specific match bool")
}
