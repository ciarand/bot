package bot

import (
	"regexp"

	"github.com/daneharrigan/hipchat"
)

type Client struct {
	conf     *config
	conn     *hipchat.Client
	handlers map[*regexp.Regexp]handler
}

type handler func(m *Message, c *Client)

func NewClient(c *config) *Client {
	return &Client{conf: c, handlers: make(map[*regexp.Regexp]handler)}
}

func (c *Client) Say(msg string) {
	// no-op
}

func (c *Client) Run() {
	// no-op
}

func (c *Client) RegisterHandler(re *regexp.Regexp, h handler) {
	c.handlers[re] = h
}

func (c *Client) handleMessage(m *Message) {
	for regex, fn := range c.handlers {
		if regex.MatchString(m.Body) {
			go fn(m, c)
			return
		}
	}
}

func (c *Client) Connect() (err error) {
	c.conn, err = hipchat.NewClient(c.conf.Username, c.conf.Password, "bot")

	return err
}
