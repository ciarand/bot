package bot

// Context is essentially a single chat room
type Context interface {
	Reply(Message)
}

// Message is the embodiment of a chat event
type Message interface {
	// Gets a "pretty" rendering of the Message
	String() string

	// Gets the body of the message
	Body() string

	// Gets the Source (usually the author) of the Message
	Source() Source

	// Gets the Context of the message
	Context() Context
}

// Responder is anything that can respond to a Message object
type Responder interface {
	Respond(Message)
}

// Source is a snapshot-in-time of a source of messages
type Source interface {
	// Gets the "pretty" name of the Source
	String() string

	// Gets an identifier string for the Source
	Id() string
}
