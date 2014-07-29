package bot

import "testing"

func TestConfFromCliFlags(t *testing.T) {
	args := []string{
		"-username", "username",
		"-room", "room_id",
		"-full_name", "full_name",
		"-mention_name", "mention_name",
	}

	conf := confFromFlags(args)

	if conf.Username != "username" ||
		conf.RoomId != "room_id" ||
		conf.FullName != "full_name" ||
		conf.MentionName != "mention_name" {
		t.Fail()
	}
}
