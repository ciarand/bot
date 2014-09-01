package bot

import (
	"errors"
	"fmt"
)

type config struct {
	Username    string
	Password    string
	RoomId      string
	FullName    string
	MentionName string
}

func (base *config) MergeWith(merges ...*config) *config {
	fin := &config{
		Username:    base.Username,
		Password:    base.Password,
		RoomId:      base.RoomId,
		FullName:    base.FullName,
		MentionName: base.MentionName,
	}

	for _, diff := range merges {
		if diff.Username != "" {
			fin.Username = diff.Username
		}

		if diff.Password != "" {
			fin.Password = diff.Password
		}

		if diff.RoomId != "" {
			fin.RoomId = diff.RoomId
		}

		if diff.FullName != "" {
			fin.FullName = diff.FullName
		}

		if diff.MentionName != "" {
			fin.MentionName = diff.MentionName
		}
	}

	return fin
}

func (c *config) validate() error {
	if c.Username == "" {
		return errors.New("missing username")
	}

	if c.Password == "" {
		return errors.New("missing password")
	}

	if c.RoomId == "" {
		return errors.New("missing room id")
	}

	if c.FullName == "" {
		return errors.New("missing full name")
	}

	if c.MentionName == "" {
		return errors.New("missing mention name")
	}

	return nil
}

func (c *config) setFromVar(key, val string) error {
	switch key {
	case "BOT_USERNAME":
		c.Username = val
		return nil
	case "BOT_PASSWORD":
		c.Password = val
		return nil
	case "BOT_ROOM_ID":
		c.RoomId = val
		return nil
	case "BOT_FULL_NAME":
		c.FullName = val
		return nil
	case "BOT_MENTION_NAME":
		c.MentionName = val
		return nil
	default:
		return fmt.Errorf("%s is not a valid variable", key)
	}
}
