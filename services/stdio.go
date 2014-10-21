package services

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ciarand/bot"
)

// StdioService is a basic Service that reads from and writes to Stdin/Stdout
type StdioService struct {
	in  io.Reader
	out io.Writer
}

// NewStdioService returns a new StdioService using os.Stdin and os.Stdout
func NewStdioService(in io.Reader, out io.Writer) *StdioService {
	if in == nil {
		in = os.Stdin
	}

	if out == nil {
		out = os.Stdout
	}

	return &StdioService{in: in, out: out}
}

// Send sends the provided message to stdin
func (s *StdioService) Send(m bot.Message) error {
	_, err := fmt.Fprintf(s.out, "%s: %s\n", m.Author.Username, m.Body)

	return err
}

// Receive blocks on an incoming message
func (s *StdioService) Receive() (bot.Message, error) {
	m := &bot.Message{Author: bot.Author{}}

	// create a new reader
	buf := bufio.NewReader(s.in)

	// username is anything up to the :
	uname, err := buf.ReadString(':')
	if err != nil {
		return *m, err
	}

	m.Author.Username = trim(uname, ": \x00")

	// body is the rest of the line, we're ignoring the error because it's no
	// longer relevant (i.e. the buf reader not even invoking the wrapped
	// io.Reader at this point)
	body, _ := buf.ReadString('\n')

	// strip spaces, newlines, null bytes
	m.Body = trim(body)

	return *m, err
}

func trim(s string, unwanted ...string) string {
	var cuts string

	if len(unwanted) == 0 {
		cuts = "\r\n\x00 \t"
	} else {
		cuts = unwanted[0]
	}

	return strings.Trim(s, cuts)
}
