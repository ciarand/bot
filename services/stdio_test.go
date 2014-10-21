package services

import (
	"bytes"
	"os"
	"testing"

	"github.com/ciarand/bot"
	"github.com/ciarand/bot/test"
)

func TestNewStdioService(t *testing.T) {
	s := NewStdioService(nil, nil)

	test.Equals(t, s.in, os.Stdin)
	test.Equals(t, s.out, os.Stdout)
}

func TestCorrectStdioRecv(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 8048))

	_, err := buf.WriteString("foobar: hello world\n")
	test.Ok(t, err)

	s := StdioService{in: buf}

	m, err := s.Receive()
	test.Ok(t, err)

	test.Equals(t, "foobar", m.Author.Username)
	test.Equals(t, "hello world", m.Body)
}

func TestStdioRecvWithUsernameError(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 8048))

	// invalid username / msg combo
	_, err := buf.WriteString("invalid username / msg combo")
	test.Ok(t, err)

	s := StdioService{in: buf}

	_, err = s.Receive()
	test.Assert(t, err != nil)
}

func bufferedReader(t *testing.T, str ...string) *bytes.Buffer {
	buf := bytes.NewBuffer(make([]byte, 8048))

	if len(str) > 0 {
		// invalid username / msg combo
		_, err := buf.WriteString(str[0])
		test.Ok(t, err)
	}

	return buf
}

func TestStdioSend(t *testing.T) {
	// https://xkcd.com/440/
	expectedMsg := "bar"
	expectedUsr := "foo"
	expected := expectedUsr + ": " + expectedMsg

	buf := bytes.NewBuffer(make([]byte, len(expected)))
	s := NewStdioService(nil, buf)
	s.Send(bot.NewMessage(bot.NewAuthor(expectedUsr), expectedMsg))

	// ignore error because
	actual, _ := buf.ReadString('\n')

	test.Equals(t, trim(expected), trim(actual))
}
