package bot

import "os"

func ConfFromEnv() config {
	return config{
		Username:    os.Getenv("HIPCHAT_USERNAME"),
		RoomId:      os.Getenv("HIPCHAT_ROOM_ID"),
		FullName:    os.Getenv("HIPCHAT_FULL_NAME"),
		MentionName: os.Getenv("HIPCHAT_MENTION_NAME"),
	}
}
