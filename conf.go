package bot

import (
	"errors"
	"fmt"
)

type config struct {
	Username    string
	RoomId      string
	FullName    string
	MentionName string
}

func (base config) mergeWith(diff config) config {
	fin := base

	if diff.Username != "" {
		fin.Username = diff.Username
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

	return fin
}

func (c *config) validate() error {
	if c.Username == "" {
		return errors.New("missing username")
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
	case "HIPCHAT_USERNAME":
		c.Username = val
		return nil
	case "HIPCHAT_ROOM_ID":
		c.RoomId = val
		return nil
	case "HIPCHAT_FULL_NAME":
		c.FullName = val
		return nil
	case "HIPCHAT_MENTION_NAME":
		c.MentionName = val
		return nil
	default:
		return fmt.Errorf("%s is not a valid variable", key)
	}
}
