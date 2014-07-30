package bot

import (
	"regexp"
)

type client struct {
	conf     *config
	handlers map[*regexp.Regexp]handler
}

type handler func(c *client)

func NewClient(c *config) *client {
	return &client{conf: c, handlers: make(map[*regexp.Regexp]handler)}
}

func (c *client) RegisterHandler(re *regexp.Regexp, h handler) {
	c.handlers[re] = h
}

func (cl *client) handleMessage(msg string) {
	for regex, fn := range cl.handlers {
		if regex.MatchString(msg) {
			go fn(cl)
			return
		}
	}
}
