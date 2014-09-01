package bot

import "os"

func ConfFromEnv() *config {
	return &config{
		Username:    os.Getenv("BOT_USERNAME"),
		RoomId:      os.Getenv("BOT_ROOM_ID"),
		FullName:    os.Getenv("BOT_FULL_NAME"),
		MentionName: os.Getenv("BOT_MENTION_NAME"),
	}
}
