package bot

// Bot is a container for a Responder and Contexts
type Bot struct {
	responder Responder
	contexts  []Context
}

// NewBot creates a new bot
func NewBot(r Responder) *Bot {
	return &Bot{responder: r}
}

// AddContext adds any number of Contexts to the Bot
func (b *Bot) AddContext(c ...Context) {
	b.contexts = append(b.contexts, c...)
}
