package bot_test

import (
	"os"
	"testing"

	bot "github.com/ciarand/hipchat-bot"
)

func TestReadsVanillaDotenvString(t *testing.T) {
	resetEnv("foo")

	bot.ParseDotEnv("foo=bar")

	if os.Getenv("foo") != "bar" {
		t.Fail()
	}
}

func TestReadsCommentsInDotenvString(t *testing.T) {
	resetEnv("foo")

	bot.ParseDotEnv("foo=bar\n#foo=baz")

	if os.Getenv("foo") != "bar" {
		t.Fail()
	}
}

func TestIgnoresBlankLinesInDotenvString(t *testing.T) {
	resetEnv("foo")

	bot.ParseDotEnv("foo=baz\n\n\n\nfoo=bar")

	if os.Getenv("foo") != "bar" {
		t.Fail()
	}
}

func TestHandlesDoubleQuotedValuesInDotenvString(t *testing.T) {
	resetEnv("foo")

	bot.ParseDotEnv(`foo="bar"`)

	if os.Getenv("foo") != "bar" {
		t.Fail()
	}
}

func TestHandlesSingleQuotedValuesInDotenvString(t *testing.T) {
	resetEnv("foo")

	bot.ParseDotEnv(`foo='bar'`)

	if os.Getenv("foo") != "bar" {
		t.Fail()
	}
}

func TestComplexExampleInDotenvString(t *testing.T) {
	resetEnv("foo")

	bot.ParseDotEnv("#comments\nfoo=bar\nbaz=\"quotes=equalssigns\"\n\n#blank lines")

	if os.Getenv("foo") != "bar" || os.Getenv("baz") != "quotes=equalssigns" {
		t.Fail()
	}
}

func TestSilentlyFailsIfNoFileExists(t *testing.T) {
	// this shouldn't fail or do anything with errors or whatever
	bot.ParseDotEnvFile("./dir_doesnt_exist")
}

func TestParsesDotEnvFileIfExists(t *testing.T) {
	resetEnv("foo", "baz")

	bot.ParseDotEnvFile("./fixtures")

	if os.Getenv("foo") != "bar" || os.Getenv("baz") != "quotes=equalssigns" {
		t.Fail()
	}
}

func resetEnv(key ...string) {
	for _, i := range key {
		os.Setenv(i, "")
	}
}
