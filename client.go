package bot

type client struct {
	conf *config
}

func NewClient(c *config) *client {
	return &client{conf: c}
}
