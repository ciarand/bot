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

func (base *config) mergeWith(diff *config) *config {
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

	fmt.Print(fin)

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
