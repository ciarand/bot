package services

import "github.com/ciarand/bot"

// Service is something that can Send and Receive messages
type Service interface {
	Send(bot.Message) error
	Receieve() (bot.Message, error)
}
