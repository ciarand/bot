package bot

import (
	"os"
	"testing"
)

func TestConfFromEnv(t *testing.T) {
	os.Setenv("HIPCHAT_USERNAME", "username")
	os.Setenv("HIPCHAT_ROOM_ID", "room_id")
	os.Setenv("HIPCHAT_FULL_NAME", "full_name")
	os.Setenv("HIPCHAT_MENTION_NAME", "mention_name")

	conf := confFromEnv()

	if conf.Username != "username" ||
		conf.RoomId != "room_id" ||
		conf.FullName != "full_name" ||
		conf.MentionName != "mention_name" {
		t.Fail()
	}
}
