package bot

import (
	"regexp"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	user := "77027_1062884"
	room := "77027_developers@conf.hipchat.com"
	full := "roboto"
	mention := "mention"

	c := NewClient(&config{
		Username:    user,
		RoomId:      room,
		FullName:    full,
		MentionName: mention,
	})

	assertSame(t, c.conf.Username, user)
	assertSame(t, c.conf.RoomId, room)
	assertSame(t, c.conf.FullName, full)
	assertSame(t, c.conf.MentionName, mention)
}

func TestHandlers(t *testing.T) {
	cli := acceptableClient()

	done := make(chan bool, 1)

	handler := func(c *client) {
		done <- true
	}

	re := regexp.MustCompile("roboto")

	cli.RegisterHandler(re, handler)
	cli.handleMessage("roboto zen")

	select {
	case <-done:
		// pass!
		return
	case <-time.After(1 * time.Millisecond): // hacky
		t.Errorf("handler not called")
	}
}

func acceptableClient() *client {
	return NewClient(&config{
		Username:    "77027_1062884",
		RoomId:      "77027_developers@conf.hipchat.com",
		FullName:    "roboto",
		MentionName: "mention",
	})
}
