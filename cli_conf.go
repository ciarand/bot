package bot

import (
	"flag"
)

func confFromFlags(args []string) *config {
	f := flag.NewFlagSet("config", flag.ContinueOnError)

	usernameFlag := f.String("username", "", "the username to connect as")
	roomFlag := f.String("room", "", "the room to connect to (or 'all')")
	fullNameFlag := f.String("full_name", "", "the bot's full name")
	mentionNameFlag := f.String("mention_name", "", "the bot's mention name")

	f.Parse(args)

	return &config{
		Username:    *usernameFlag,
		RoomId:      *roomFlag,
		FullName:    *fullNameFlag,
		MentionName: *mentionNameFlag,
	}
}
