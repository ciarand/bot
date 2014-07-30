package bot

import (
	"os"
	"testing"
)

var dotconfPerms = []struct {
	str    string
	expect string
	msg    string
}{
	{"HIPCHAT_USERNAME=foo", "foo", "vanilla dotconf line"},
	{"HIPCHAT_USERNAME=foo\n#HIPCHAT_USERNAME=bar", "foo", "comments"},
	{"HIPCHAT_USERNAME=bar\n\n\nHIPCHAT_USERNAME=foo", "foo", "overriding subsequent lines"},
	{"HIPCHAT_USERNAME=\"foo\"", "foo", "double quotes in value"},
	{"HIPCHAT_USERNAME='foo'", "foo", "single quotes in value"},
	{"HIPCHAT_USERNAME\nHIPCHAT_USERNAME=foo", "foo", "broken lines"},
	{"HIPCHAT_USERNAME='foo=bar'", "foo=bar", "equals in the value"},
}

func TestDotconfPerms(t *testing.T) {
	for _, tt := range dotconfPerms {
		c := parseDotConf(tt.str)

		if c.Username != tt.expect {
			t.Errorf(tt.msg)
		}
	}
}

func TestSilentlyFailsIfNoFileExists(t *testing.T) {
	// this shouldn't fail or do anything with errors or whatever
	ConfFromDotConf("./dir_doesnt_exist")
}

func TestParsesdotconfFileIfExists(t *testing.T) {
	resetEnv("foo", "baz")

	c := ConfFromDotConf("./fixtures")

	assertSame(t, c.Username, "my cool name")
	assertSame(t, c.MentionName, "nope not here#\"")
}

func resetEnv(key ...string) {
	for _, i := range key {
		os.Setenv(i, "")
	}
}