package main

import (
	"github.com/ciarand/bot"
	"github.com/ciarand/bot/services"
)

func main() {
	s := services.NewStdioService()

	s.Send(bot.NewMessage(bot.NewAuthor("ciarand"), "hello world!"))
}
