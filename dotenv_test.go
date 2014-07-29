package bot

import (
	"os"
	"testing"
)

func TestReadsVanillaDotenvString(t *testing.T) {
	resetEnv("foo")

	parseDotEnv("foo=bar")

	if os.Getenv("foo") != "bar" {
		t.Fail()
	}
}

func TestReadsCommentsInDotenvString(t *testing.T) {
	resetEnv("foo")

	parseDotEnv("foo=bar\n#foo=baz")

	if os.Getenv("foo") != "bar" {
		t.Fail()
	}
}

func TestIgnoresBlankLinesInDotenvString(t *testing.T) {
	resetEnv("foo")

	parseDotEnv("foo=baz\n\n\n\nfoo=bar")

	if os.Getenv("foo") != "bar" {
		t.Fail()
	}
}

func TestHandlesDoubleQuotedValuesInDotenvString(t *testing.T) {
	resetEnv("foo")

	parseDotEnv(`foo="bar"`)

	if os.Getenv("foo") != "bar" {
		t.Fail()
	}
}

func TestHandlesSingleQuotedValuesInDotenvString(t *testing.T) {
	resetEnv("foo")

	parseDotEnv(`foo='bar'`)

	if os.Getenv("foo") != "bar" {
		t.Fail()
	}
}

func TestIgnoresBrokenLinesInDotenvString(t *testing.T) {
	resetEnv("foo")

	parseDotEnv("foo=\nbar")

	if os.Getenv("foo") == "bar" {
		t.Fail()
	}
}

func TestComplexExampleInDotenvString(t *testing.T) {
	resetEnv("foo")

	parseDotEnv("#comments\nfoo=bar\nbaz=\"quotes=equalssigns\"\n\n#blank lines")

	if os.Getenv("foo") != "bar" || os.Getenv("baz") != "quotes=equalssigns" {
		t.Fail()
	}
}

func TestSilentlyFailsIfNoFileExists(t *testing.T) {
	// this shouldn't fail or do anything with errors or whatever
	ParseDotEnvFile("./dir_doesnt_exist")
}

func TestParsesDotEnvFileIfExists(t *testing.T) {
	resetEnv("foo", "baz")

	ParseDotEnvFile("./fixtures")

	if os.Getenv("foo") != "bar" || os.Getenv("baz") != "quotes=equalssigns" {
		t.Fail()
	}
}

func resetEnv(key ...string) {
	for _, i := range key {
		os.Setenv(i, "")
	}
}
