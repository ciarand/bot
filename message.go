package bot

import (
	"strings"
	"time"

	"github.com/daneharrigan/hipchat"
)

type Message interface {
	From() string
	Body() string
	Timestamp() time.Time
}

type message struct {
	from      string
	body      string
	timestamp time.Time
}

func messageFromHipchatLib(msg *hipchat.Message) *message {
	return &message{from: msg.From, body: msg.Body, timestamp: time.Now()}
}

func (m *message) From() string {
	return strings.TrimSpace(m.from)
}

func (m *message) Body() string {
	return m.body
}

func (m *message) Timestamp() time.Time {
	return m.timestamp
}
