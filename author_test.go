package bot

import (
	"testing"

	"github.com/ciarand/bot/test"
)

func TestNewAuthor(t *testing.T) {
	test.Equals(t,
		Author{"foobar"},
		NewAuthor("foobar"))
}
