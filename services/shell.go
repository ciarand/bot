package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ciarand/bot"
)

// StdioService is a basic Service that reads from and writes to Stdin/Stdout
type StdioService struct {
	in  *os.File
	out *os.File
}

// NewStdioService returns a new StdioService using os.Stdin and os.Stdout
func NewStdioService() *StdioService {
	return &StdioService{in: os.Stdin, out: os.Stdout}
}

// Send sends the provided message to stdin
func (s *StdioService) Send(m bot.Message) error {
	_, err := fmt.Fprintf(s.out, "%s: %s\n", m.Author.Username, m.Body)

	return err
}

// Receive blocks on an incoming message
func (s *StdioService) Receive() (bot.Message, error) {
	var err error
	var str string

	m := &bot.Message{Author: bot.Author{}}

	// create a new reader
	buf := bufio.NewReader(s.in)

	// username is anything up to the :
	if str, err = buf.ReadString(':'); err != nil {
		return *m, err
	}

	m.Author.Username = strings.Trim(str, ":")

	// body is the rest of the line
	if str, err = buf.ReadString('\n'); err != nil {
		return *m, err
	}

	m.Body = strings.TrimSpace(str)

	return *m, err
}
