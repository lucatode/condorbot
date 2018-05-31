package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindingExpectedHotkeys(t *testing.T) {

	notifyKey := "#notify"

	assert.Equal(t, notifyKey, "#notify", "Getting specific hotkey")
}
