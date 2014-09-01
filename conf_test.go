package bot

import (
	"errors"
	"strings"
	"testing"
)

func TestMergeConfig(t *testing.T) {
	foo := &config{
		Username:    "foo",
		Password:    "supersecret",
		RoomId:      "0000000",
		FullName:    "name1",
		MentionName: "mentionname1",
	}

	bar := &config{
		Username:    "bar",
		RoomId:      "1111111",
		FullName:    "name2",
		MentionName: "mentionname2",
	}

	final := foo.MergeWith(bar)

	assertSame(t, final.Username, "bar")
	assertSame(t, final.RoomId, "1111111")
	assertSame(t, final.FullName, "name2")
	assertSame(t, final.MentionName, "mentionname2")
	assertSame(t, final.Password, "supersecret")

	ulti := bar.MergeWith(foo)

	assertSame(t, ulti.Username, "foo")
	assertSame(t, ulti.RoomId, "0000000")
	assertSame(t, ulti.FullName, "name1")
	assertSame(t, ulti.MentionName, "mentionname1")
	assertSame(t, ulti.Password, "supersecret")
}

var confPerms = []struct {
	c      config
	prefix string
}{
	{config{
		Username:    "bar",
		Password:    "secret",
		RoomId:      "1111111",
		FullName:    "",
		MentionName: "mentionname2",
	}, "missing full name"},

	{config{
		Username:    "",
		Password:    "secret",
		RoomId:      "1111111",
		FullName:    "",
		MentionName: "mentionname2",
	}, "missing username"},

	{config{
		Username:    "bar",
		Password:    "secret",
		RoomId:      "",
		FullName:    "",
		MentionName: "mentionname2",
	}, "missing room id"},

	{config{
		Username:    "bar",
		Password:    "secret",
		RoomId:      "1111111",
		FullName:    "asjdksadsa",
		MentionName: "",
	}, "missing mention name"},

	{config{
		Username:    "bar",
		Password:    "",
		RoomId:      "1111111",
		FullName:    "asjdksadsa",
		MentionName: "blah",
	}, "missing password"},
}

func TestValidateConfig(t *testing.T) {
	for _, p := range confPerms {
		err := p.c.validate()

		if err == nil {
			t.Errorf("no error returned for %s", p.prefix)
		} else if !strings.HasPrefix(err.Error(), p.prefix) {
			t.Errorf("missing prefix: %s", p.prefix)
		}
	}

	// finally validate a correct one
	correct := &config{
		Username:    "foo",
		Password:    "secret",
		RoomId:      "0000000",
		FullName:    "name1",
		MentionName: "mentionname1",
	}

	err := correct.validate()
	if err != nil {
		t.Errorf("failed to validate correct config")
	}
}

var setFromVarPerms = []struct {
	key string
	err error
	val string
}{
	{"BOT_USERNAME", nil, "foo"},
	{"BOT_USERNAME", nil, "bar"},
	{"BOT_PASSWORD", nil, "foo"},
	{"BOT_ROOM_ID", nil, "foo"},
	{"BOT_FULL_NAME", nil, "foo"},
	{"BOT_MENTION_NAME", nil, "foo"},
	{"FOOBAR", errors.New("FOOBAR is not a valid variable"), ""},
}

func TestSetFromVar(t *testing.T) {
	c := &config{}

	for _, tt := range setFromVarPerms {
		err := c.setFromVar(tt.key, tt.val)

		if tt.err != nil {
			if err == nil {
				t.Errorf("expected error '%s', but none found", tt.err)
				continue
			}

			if err.Error() != tt.err.Error() {
				t.Errorf("error messages don't match. expected %s, got %s",
					tt.err.Error(), err.Error())
				continue
			}
		} else {
			switch tt.key {
			case "BOT_USERNAME":
				assertSame(t, c.Username, tt.val)
				break
			case "BOT_PASSWORD":
				assertSame(t, c.Password, tt.val)
				break
			case "BOT_ROOM_ID":
				assertSame(t, c.RoomId, tt.val)
				break
			case "BOT_FULL_NAME":
				assertSame(t, c.FullName, tt.val)
				break
			case "BOT_MENTION_NAME":
				assertSame(t, c.MentionName, tt.val)
				break
			}
		}
	}
}
