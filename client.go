package bot

import (
	"errors"
	"log"
	"regexp"
	"time"

	"github.com/daneharrigan/hipchat"
)

type Client struct {
	conf     *config
	conn     *hipchat.Client
	handlers map[*regexp.Regexp]handler
}

type handler func(m Message, c *Client)

func NewClient(c *config) *Client {
	return &Client{conf: c, handlers: make(map[*regexp.Regexp]handler)}
}

func (c *Client) Send(msg string) {
	log.Printf("SEND: [%s] %s: %s",
		time.Now().Format("Jan 2 15:04:05"), c.conf.FullName, msg)

	c.conn.Say(c.conf.RoomId, c.conf.FullName, msg)
}

func (c *Client) Run() error {
	if c.conn == nil {
		return errors.New("not connected to the HipChat server")
	}

	c.conn.Join(c.conf.RoomId, c.conf.FullName)

	selfPattern := regexp.MustCompile(`/` + c.conf.FullName + `$`)

	for msg := range c.conn.Messages() {
		m := messageFromHipchatLib(msg)

		// if we said it, ignore it
		if selfPattern.MatchString(m.From()) {
			continue
		}

		log.Printf("RECV: [%s] %s: %s",
			m.Timestamp().Format("Jan 2 15:04:05"), m.From(), m.Body())

		go c.handleMessage(m)
	}

	return nil
}

func (c *Client) RegisterHandler(re *regexp.Regexp, h handler) {
	c.handlers[re] = h
}

func (c *Client) handleMessage(m Message) {
	for regex, fn := range c.handlers {
		if regex.MatchString(m.Body()) {
			fn(m, c)
			return
		}
	}
}

func (c *Client) Connect() error {
	var err error

	if err = c.conf.validate(); err != nil {
		return err
	}

	if c.conn, err = hipchat.NewClient(c.conf.Username, c.conf.Password, "bot"); err != nil {
		return err
	}

	c.conn.Status("chat")

	go c.conn.KeepAlive()

	return nil
}
