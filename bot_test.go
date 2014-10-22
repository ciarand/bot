package bot

import (
	"testing"

	"github.com/ciarand/bot/test"
)

func TestAddContext(t *testing.T) {
	b := NewBot(nullResponder{})

	b.AddContext(nullContext{})

	test.Equals(t, 1, len(b.contexts))
}

func TestNewBotAddsResponder(t *testing.T) {
	b := NewBot(nullResponder{})

	test.Equals(t, nullResponder{}, b.responder)
}
