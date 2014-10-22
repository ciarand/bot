package bot

import (
	"testing"

	"github.com/ciarand/bot/test"
)

type channelResponder struct {
	c chan Message
}

func newChannelResponder() *channelResponder {
	return &channelResponder{c: make(chan Message)}
}

func (c *channelResponder) Respond(m Message) {
	c.c <- m
}

type nullContext struct{}

func (n nullContext) Reply(m Message) {
}

func TestAddContext(t *testing.T) {
	b := NewBot(newChannelResponder())

	b.AddContext(nullContext{})

	test.Equals(t, 1, len(b.contexts))
}
