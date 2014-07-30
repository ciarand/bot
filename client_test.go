package bot

import (
	"testing"
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
