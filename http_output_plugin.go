package bot

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ciarand/bot"
)

func HttpOutputPlugin(url string) handler {

	return func(msg bot.Message, c *bot.Client) {
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
