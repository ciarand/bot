package bot

import "testing"

func TestConfFromCliFlags(t *testing.T) {
	args := []string{
		"-username", "username",
		"-room", "room_id",
		"-password", "secret",
		"-full_name", "full_name",
		"-mention_name", "mention_name",
	}

	conf := ConfFromFlags(args)

	assertSame(t, conf.Username, "username")
	assertSame(t, conf.Password, "secret")
	assertSame(t, conf.RoomId, "room_id")
	assertSame(t, conf.FullName, "full_name")
	assertSame(t, conf.MentionName, "mention_name")
}
