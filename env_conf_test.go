package bot

import (
	"os"
	"testing"
)

func TestConfFromEnv(t *testing.T) {
	os.Setenv("BOT_USERNAME", "username")
	os.Setenv("BOT_ROOM_ID", "room_id")
	os.Setenv("BOT_FULL_NAME", "full_name")
	os.Setenv("BOT_MENTION_NAME", "mention_name")

	conf := ConfFromEnv()

	if conf.Username != "username" ||
		conf.RoomId != "room_id" ||
		conf.FullName != "full_name" ||
		conf.MentionName != "mention_name" {
		t.Fail()
	}
}
