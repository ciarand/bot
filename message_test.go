package bot

import (
	"testing"

	"github.com/ciarand/bot/test"
)

func TestNewMessage(t *testing.T) {
	test.Equals(t,
		NewMessage(NewAuthor("foobar"), "hello world"),
		Message{Author: Author{"foobar"}, Body: "hello world"})
}
