package main

import (
	"os"

	"github.com/ciarand/bot"
	"github.com/ciarand/bot/services"
)

func main() {
	s := services.NewStdioService(os.Stdin, os.Stdout)

	s.Send(bot.NewMessage(bot.NewAuthor("ciarand"), "hello world!"))
}
