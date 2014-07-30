package bot

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpOutputPlugin(url string) handler {

	return func(msg Message, c *Client) {
		resp, err := http.Get(url)
		if err != nil {
			c.Send(fmt.Sprintf("couldn't retrieve page: %s", err.Error()))
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.Send(fmt.Sprintf("couldn't read the response: %s", err.Error()))
			return
		}

		c.Send(string(body))
	}
}
