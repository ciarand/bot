package bot

// Author is the entity responsible for sending messages
type Author struct {
	Username string
}

// NewAuthor creates a new Author based on the provided username
func NewAuthor(username string) Author {
	return Author{username}
}
