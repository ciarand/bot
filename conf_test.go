package bot

import (
	"strings"
	"testing"
)

func TestMergeConfig(t *testing.T) {
	first := &config{
		Username:    "foo",
		RoomId:      "0000000",
		FullName:    "name1",
		MentionName: "mentionname1",
	}

	second := &config{
		Username:    "bar",
		RoomId:      "1111111",
		FullName:    "",
		MentionName: "mentionname2",
	}

	final := first.mergeWith(second)

	if final.Username != "bar" ||
		final.RoomId != "1111111" ||
		final.FullName != "name1" ||
		final.MentionName != "mentionname2" {
		t.Fail()
	}
}

var perms = []struct {
	c      config
	prefix string
}{
	{config{
		Username:    "bar",
		RoomId:      "1111111",
		FullName:    "",
		MentionName: "mentionname2",
	}, "missing full name"},

	{config{
		Username:    "",
		RoomId:      "1111111",
		FullName:    "",
		MentionName: "mentionname2",
	}, "missing username"},

	{config{
		Username:    "bar",
		RoomId:      "",
		FullName:    "",
		MentionName: "mentionname2",
	}, "missing room id"},

	{config{
		Username:    "bar",
		RoomId:      "1111111",
		FullName:    "asjdksadsa",
		MentionName: "",
	}, "missing mention name"},
}

func testValidateConfig(t *testing.T) {
	for _, p := range perms {
		err := p.c.validate()

		if err == nil {
			t.Errorf("no error returned for %s", p.prefix)
		}

		if !strings.HasPrefix(err.Error(), p.prefix) {
			t.Errorf("missing prefix")
		}
	}
}
