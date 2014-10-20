package bot

// Message is the atomic conversation unit
type Message struct {
	Body   string
	Author Author
}

// NewMessage creates a new message from the provided string
func NewMessage(u Author, s string) Message {
	return Message{Author: u, Body: s}
}
